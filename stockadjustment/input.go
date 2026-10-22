package stockadjustment

type CreateStockAdjustmentInput struct {
	ProductID int    `json:"product_id" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=increase decrease"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
	Reason    string `json:"reason" binding:"required"`
	Reference string `json:"reference"`
	CreatedBy int    `json:"-"`
}

type GetStockAdjustmentInput struct {
	ID int `uri:"id" binding:"required"`
}
