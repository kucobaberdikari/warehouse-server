package mutation

import (
	"management-warehouse/product"
	"management-warehouse/user"
)

type GetMutationDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMutationInput struct {
	Type     string          `json:"type" binding:"required"`
	Quantity int             `json:"quantity" binding:"required"`
	Note     string          `json:"note" binding:"required"`
	Product  product.Product `binding:"required"`
	User     user.User       `json:"-"`
}

type FormCreateMutationInput struct {
	Type     string            `json:"type" binding:"required"`
	Quantity int               `json:"quantity" binding:"required"`
	Note     string            `json:"note" binding:"required"`
	Product  []product.Product `binding:"required"`
	User     []user.User       `binding:"required"`
	Error    error
}
