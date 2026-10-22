package stockadjustment

import "time"

type StockAdjustmentFormat struct {
	ID          int            `json:"id"`
	ProductID   int            `json:"product_id"`
	ProductName string         `json:"product_name,omitempty"`
	Type        AdjustmentType `json:"type"`
	Quantity    int            `json:"quantity"`
	Reason      string         `json:"reason"`
	Reference   string         `json:"reference,omitempty"`
	CreatedBy   int            `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func StockAdjustmentsFormat(stockadjustment []StockAdjustment) []StockAdjustmentFormat {
	result := []StockAdjustmentFormat{}

	for _, stockadjustment := range stockadjustment {
		result = append(result, StockAdjustmentFormat(stockadjustment))
	}

	return result
}
