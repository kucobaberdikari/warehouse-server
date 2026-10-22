package mutation

import "errors"

type Service interface {
	GetMutation() ([]Mutation, error)
	GetMutationByID(input GetMutationDetailInput) (Mutation, error)
	CreateMutation(input CreateMutationInput) (Mutation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMutation() ([]Mutation, error) {
	mutation, err := s.repository.findAll()
	if err != nil {
		return mutation, err
	}

	return mutation, nil
}

func (s *service) GetMutationByID(input GetMutationDetailInput) (Mutation, error) {
	mutation, err := s.repository.FindByID(input.ID)
	if err != nil {
		return mutation, err
	}

	return mutation, nil
}

func (s *service) CreateMutation(input CreateMutationInput) (Mutation, error) {
	mutation := Mutation{}

	mutation.Type = MutationType(input.Type)
	if mutation.Type == "in" {
		mutation.Type = MutationTypeInbound
	}
	if mutation.Type == "out" {
		mutation.Type = MutationTypeOutbound
	}

	if mutation.Type != MutationTypeInbound && mutation.Type != MutationTypeOutbound {
		return mutation, errors.New("mutation type must be inbound or outbound")
	}
	if input.Quantity <= 0 {
		return mutation, errors.New("quantity must be greater than zero")
	}
	if input.Product.ID == 0 {
		return mutation, errors.New("product id is required")
	}
	if input.User.ID == 0 {
		return mutation, errors.New("user id is required")
	}

	mutation.ProductID = input.Product.ID
	mutation.UserID = input.User.ID
	mutation.Quantity = input.Quantity
	mutation.Note = input.Note

	newMutation, err := s.repository.Save(mutation)
	if err != nil {
		return newMutation, err
	}

	return newMutation, nil
}
