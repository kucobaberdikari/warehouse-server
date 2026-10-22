package stockadjustment

import (
	"errors"
	"time"

	"management-warehouse/product"

	"gorm.io/gorm"
)

type Repository interface {
	Save(adjustment StockAdjustment) (StockAdjustment, error)
	FindAll() ([]StockAdjustment, error)
	FindByID(id int) (StockAdjustment, error)
	FindByProductID(productID int) ([]StockAdjustment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(adjustment StockAdjustment) (StockAdjustment, error) {
	if adjustment.ProductID <= 0 {
		return adjustment, ErrProductIDRequired
	}
	if adjustment.Quantity <= 0 {
		return adjustment, ErrInvalidQuantity
	}

	stockChange := adjustment.Quantity
	switch adjustment.Type {
	case AdjustmentIncrease:
	case AdjustmentDecrease:
		stockChange = -adjustment.Quantity
	default:
		return adjustment, ErrInvalidAdjustmentType
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		stockUpdate := tx.Model(&product.Product{}).Where("id = ?", adjustment.ProductID)
		if adjustment.Type == AdjustmentDecrease {
			stockUpdate = stockUpdate.Where("stock >= ?", adjustment.Quantity)
		}

		result := stockUpdate.Updates(map[string]interface{}{
			"stock":      gorm.Expr("stock + ?", stockChange),
			"updated_at": time.Now(),
		})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			var productCount int64
			if err := tx.Model(&product.Product{}).Where("id = ?", adjustment.ProductID).Count(&productCount).Error; err != nil {
				return err
			}
			if productCount == 0 {
				return errors.New("product not found")
			}
			return errors.New("insufficient stock")
		}

		if err := tx.Create(&adjustment).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return adjustment, err
	}
	return adjustment, nil
}

func (r *repository) FindAll() ([]StockAdjustment, error) {
	var adjustments []StockAdjustment
	if err := r.db.Order("created_at desc").Find(&adjustments).Error; err != nil {
		return nil, err
	}
	return adjustments, nil
}

func (r *repository) FindByID(id int) (StockAdjustment, error) {
	var adjustment StockAdjustment
	if err := r.db.Where("id = ?", id).First(&adjustment).Error; err != nil {
		return adjustment, err
	}
	return adjustment, nil
}

func (r *repository) FindByProductID(productID int) ([]StockAdjustment, error) {
	var adjustments []StockAdjustment
	if err := r.db.Where("product_id = ?", productID).Order("created_at desc").Find(&adjustments).Error; err != nil {
		return nil, err
	}
	return adjustments, nil
}
