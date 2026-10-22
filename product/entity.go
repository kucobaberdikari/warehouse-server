package product

import (
	"management-warehouse/category"
	"management-warehouse/supplier"
	"time"
)

type Product struct {
	ID         int               `gorm:"primaryKey"`
	SKU        string            `gorm:"not null,unique"`
	Name       string            `gorm:"not null,unique"`
	Stock      int               `gorm:"not null" `
	MinStock   int               `gorm:"not null"`
	Unit       string            `gorm:"not null"`
	CategoryID int               `gorm:"column:category_id" json:"category_id"`
	Category   category.Category `gorm:"foreignKey:CategoryID;references:ID"`
	SupplierID int               `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier   supplier.Supplier `gorm:"foreignKey:SupplierID;references:ID"`
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
