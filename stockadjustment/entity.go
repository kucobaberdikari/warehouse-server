package stockadjustment

import "time"

type AdjustmentType string

const (
	AdjustmentIncrease AdjustmentType = "increase"
	AdjustmentDecrease AdjustmentType = "decrease"
)

type StockAdjustment struct {
	ID          int
	ProductID   int
	ProductName string
	Type        AdjustmentType
	Quantity    int
	Reason      string
	Reference   string
	CreatedBy   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
