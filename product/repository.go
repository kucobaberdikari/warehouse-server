package product

import (
	"management-warehouse/category"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Product, error)
	FindBySupplierID(supplierID int) ([]Product, error)
	FindByID(ID int) (Product, error)
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	DeleteProduct(ID int) (Product, error)
	SearchProduct(Name, IDSupplier, SKU, Category string) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var product []Product

	err := r.db.Preload("Category").Preload("Supplier").Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindBySupplierID(supplierID int) ([]Product, error) {
	var product []Product

	err := r.db.Where("supplier_id = ?", supplierID).Preload("Category").Preload("Supplier").Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product

	err := r.db.Preload("Category").Preload("Supplier").Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
func (r *repository) DeleteProduct(ID int) (Product, error) {
	var product Product
	err := r.db.Where("id = ?", ID).Find(&product).Delete(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) SearchProduct(Name, IDSupplier, SKU, Category string) ([]Product, error) {
	var product []Product

	filter := r.db
	filters := false
	if Name != "" {
		pattern := "%" + strings.ToLower(Name) + "%"
		filter = filter.Where("LOWER(name) LIKE ?", pattern)
		filters = true
	}

	if SKU != "" {
		pattern := "%" + strings.ToLower(SKU) + "%"
		skuFilter := r.db.Where("LOWER(sku) LIKE ?", pattern)
		if filters {
			filter = filter.Or(skuFilter)
		} else {
			filter = skuFilter
		}
		filters = true
	}

	if Category != "" {
		pattern := "%" + strings.ToLower(Category) + "%"
		categoryIDs := r.db.Model(&category.Category{}).
			Select("id").
			Where("LOWER(name) LIKE ?", pattern)
		categoryFilter := r.db.Where("category_id IN (?)", categoryIDs)
		if categoryID, err := strconv.Atoi(Category); err == nil {
			categoryFilter = categoryFilter.Or("category_id = ?", categoryID)
		}
		if filters {
			filter = filter.Or(categoryFilter)
		} else {
			filter = categoryFilter
		}
		filters = true
	}

	if supplierID, err := strconv.Atoi(IDSupplier); IDSupplier != "" && err == nil {
		supplierFilter := r.db.Where("supplier_id = ?", supplierID)
		if filters {
			filter = filter.Or(supplierFilter)
		} else {
			filter = supplierFilter
		}
		filters = true
	}

	query := r.db.Model(&Product{})
	if filters {
		query = query.Where(filter)
	}
	err := query.Preload("Category").Preload("Supplier").Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
