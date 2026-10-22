package mutation

import (
	"testing"

	"management-warehouse/product"
	"management-warehouse/user"
)

type stubMutationRepository struct {
	saved Mutation
}

func (s *stubMutationRepository) findAll() ([]Mutation, error) {
	return nil, nil
}

func (s *stubMutationRepository) FindByID(ID int) (Mutation, error) {
	return Mutation{}, nil
}

func (s *stubMutationRepository) Save(m Mutation) (Mutation, error) {
	s.saved = m
	return m, nil
}

func TestCreateMutationSetsForeignKeys(t *testing.T) {
	repo := &stubMutationRepository{}
	service := NewService(repo)

	input := CreateMutationInput{
		Type:     "inbound",
		Quantity: 5,
		Note:     "stok masuk",
		Product:  product.Product{ID: 7},
		User:     user.User{ID: 9},
	}

	got, err := service.CreateMutation(input)
	if err != nil {
		t.Fatalf("CreateMutation returned error: %v", err)
	}

	if got.ProductID != input.Product.ID {
		t.Fatalf("expected ProductID %d, got %d", input.Product.ID, got.ProductID)
	}

	if got.UserID != input.User.ID {
		t.Fatalf("expected UserID %d, got %d", input.User.ID, got.UserID)
	}
}
