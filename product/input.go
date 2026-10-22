package product

import (
	"management-warehouse/category"
	"management-warehouse/supplier"
)

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetSupplierProductDetailInput struct {
	SupplierID int `uri:"supplier_id" binding:"required"`
}

type CreateProductInput struct {
	// ID         int    `json:"id" binding:"required"`
	SKU        string            `json:"sku" binding:"required" gorm:"unique"`
	Name       string            `json:"name" binding:"required" gorm:"unique"`
	Stock      int               `json:"stock" binding:"required"`
	MinStock   int               `json:"min_stock" binding:"required"`
	Unit       string            `json:"unit" binding:"required"`
	SupplierID int               `json:"supplier_id" binding:"required"`
	Image      string            `json:"image"`
	Category   category.Category `binding:"required"`
}

type FormCreateProductInput struct {
	SKU      string `form:"sku" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Stock    int    `form:"stock" binding:"required"`
	MinStock int    `form:"min_stock" binding:"required"`
	Unit     string `form:"unit" binding:"required"`
	// SupplierID int    `json:"supplier_id"`
	Image    string              `json:"image"`
	Supplier []supplier.Supplier `binding:"required"`
	Category []category.Category `binding:"required"`
	Error    error
}

type FormUpdateProductInput struct {
	// ID         int    `binding:"required"`
	SKU        string `json:"sku" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	MinStock   int    `json:"min_stock" binding:"required"`
	Unit       string `json:"unit" binding:"required"`
	SupplierID int    `json:"supplier_id" binding:"required"`
	Image      string `json:"image" binding:"required"`
	Error      error
}

type GetSearchProduct struct {
	Name string `uri:"name" binding:"required"`
}

type GetProductSearchInput struct {
	Name       string `form:"name"`
	SupplierID string `form:"supplier_id"`
	CategoryID string `form:"category_id"`
	SKU        string `form:"sku"`
}
