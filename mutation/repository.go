package mutation

import (
	"errors"
	"time"

	"management-warehouse/product"

	"gorm.io/gorm"
)

type Repository interface {
	findAll() ([]Mutation, error)
	FindByID(ID int) (Mutation, error)
	Save(mutation Mutation) (Mutation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) findAll() ([]Mutation, error) {
	var mutation []Mutation

	err := r.db.
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Supplier").
		Preload("User").
		Find(&mutation).Error
	if err != nil {
		return mutation, err
	}

	return mutation, nil
}

func (r *repository) FindByID(ID int) (Mutation, error) {
	var mutation Mutation

	err := r.db.
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Supplier").
		Preload("User").
		Where("id = ?", ID).
		First(&mutation).Error
	if err != nil {
		return mutation, err
	}

	return mutation, nil
}

func (r *repository) Save(mutation Mutation) (Mutation, error) {
	if mutation.Quantity <= 0 {
		return mutation, errors.New("quantity must be greater than zero")
	}

	stockChange := mutation.Quantity
	switch mutation.Type {
	case "inbound":
	case "outbound":
		stockChange = -mutation.Quantity
	default:
		return mutation, errors.New("mutation type must be inbound or outbound")
	}

	productID := mutation.ProductID
	if productID == 0 {
		productID = mutation.Product.ID
	}
	mutation.ProductID = productID

	err := r.db.Transaction(func(tx *gorm.DB) error {
		stockUpdate := tx.Model(&product.Product{}).Where("id = ?", productID)
		if mutation.Type == "outbound" {
			stockUpdate = stockUpdate.Where("stock >= ?", mutation.Quantity)
		}

		result := stockUpdate.Updates(map[string]interface{}{
			"stock":      gorm.Expr("stock + ?", stockChange),
			"updated_at": time.Now(),
		})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			if mutation.Type == "outbound" {
				var count int64
				if err := tx.Model(&product.Product{}).Where("id = ?", productID).Count(&count).Error; err != nil {
					return err
				}
				if count > 0 {
					return errors.New("insufficient stock")
				}
			}
			return errors.New("product not found")
		}

		if err := tx.Omit("Product", "User").Create(&mutation).Error; err != nil {
			return err
		}

		return tx.
			Preload("Product").
			Preload("Product.Category").
			Preload("Product.Supplier").
			Preload("User").
			First(&mutation, mutation.ID).Error
	})
	if err != nil {
		return mutation, err
	}

	return mutation, nil
}
