package stockadjustment

import "errors"

type Service interface {
	Create(input CreateStockAdjustmentInput) (StockAdjustment, error)
	GetAll() ([]StockAdjustment, error)
	GetByID(input GetStockAdjustmentInput) (StockAdjustment, error)
	GetByProductID(productID int) ([]StockAdjustment, error)
}

type service struct {
	repository Repository
}

var (
	ErrInvalidAdjustmentType = errors.New("invalid adjustment type")
	ErrInvalidQuantity       = errors.New("quantity must be greater than zero")
	ErrProductIDRequired     = errors.New("product id is required")
)

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateStockAdjustmentInput) (StockAdjustment, error) {
	adjustment := StockAdjustment{
		ProductID: input.ProductID,
		Reason:    input.Reason,
		Reference: input.Reference,
		CreatedBy: input.CreatedBy,
		Quantity:  input.Quantity,
	}
	if input.ProductID <= 0 {
		return adjustment, ErrProductIDRequired
	}
	if input.Quantity <= 0 {
		return adjustment, ErrInvalidQuantity
	}

	switch input.Type {
	case "increase":
		adjustment.Type = AdjustmentIncrease
	case "decrease":
		adjustment.Type = AdjustmentDecrease
	default:
		return adjustment, ErrInvalidAdjustmentType
	}

	return s.repository.Save(adjustment)
}

func (s *service) GetAll() ([]StockAdjustment, error) {
	return s.repository.FindAll()
}

func (s *service) GetByID(input GetStockAdjustmentInput) (StockAdjustment, error) {
	return s.repository.FindByID(input.ID)
}

func (s *service) GetByProductID(productID int) ([]StockAdjustment, error) {
	return s.repository.FindByProductID(productID)
}
