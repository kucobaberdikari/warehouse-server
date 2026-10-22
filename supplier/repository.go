package supplier

import (
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Save(supplier Supplier) (Supplier, error)
	FindByID(ID int) (Supplier, error)
	FindAll() ([]Supplier, error)
	Update(supplier Supplier) (Supplier, error)
	SearchSupplier(Name string) ([]Supplier, error)
	DeleteSupplier(ID int) (Supplier, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(supplier Supplier) (Supplier, error) {
	err := r.db.Create(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FindByID(ID int) (Supplier, error) {
	var supplier Supplier

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FindAll() ([]Supplier, error) {
	var supplier []Supplier

	err := r.db.Find("Supplier").Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) Update(supplier Supplier) (Supplier, error) {
	err := r.db.Save(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil

}

func (r *repository) SearchSupplier(Name string) ([]Supplier, error) {
	var supplier []Supplier

	filter := r.db
	filters := false
	if Name != "" {
		pattern := "%" + strings.ToLower(Name) + "%"
		filter = filter.Where("LOWER(name) LIKE ?", pattern)
		filters = true
	}
	query := r.db.Model(&Supplier{})
	if filters {
		query = query.Where(filter)
	}
	err := query.Find(&supplier).Error
	if err != nil {
		return supplier, err
	}
	return supplier, nil

}

func (r *repository) DeleteSupplier(ID int) (Supplier, error) {
	var supplier Supplier
	err := r.db.Where("id +?", ID).Find(&supplier).Delete(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}
