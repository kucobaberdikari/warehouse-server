package supplier

import (
	"errors"
)

type Service interface {
	RegisterSupplier(input RegisterSupplierInput) (Supplier, error)
	GetSupplierByID(input GetSupplierByID) (Supplier, error)
	FindAll() ([]Supplier, error)
	Updatesupplier(input GetSupplierByID, inputData FormUpdateSupplierInput) (Supplier, error)
	GetSupplierDetail(input GetSupplierByID) (Supplier, error)
	SearchSUpplier(input GetSupplierSearchInput) ([]Supplier, error)
	DeleteSupplier(input GetSupplierByID) (Supplier, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterSupplier(input RegisterSupplierInput) (Supplier, error) {

	supplier := Supplier{}
	supplier.Name = input.Name
	supplier.ContactPerson = input.ContactPerson
	supplier.Phone = input.Phone
	supplier.Address = input.Address
	newSupplier, err := s.repository.Save(supplier)
	if err != nil {
		return newSupplier, err
	}

	return newSupplier, nil
}

func (s *service) GetSupplierByID(input GetSupplierByID) (Supplier, error) {
	supplier, err := s.repository.FindByID(input.ID)
	if err != nil {
		return supplier, err
	}

	if supplier.ID == 0 {
		return supplier, errors.New("No supplier Found")
	}

	return supplier, nil
}

func (s *service) Updatesupplier(input GetSupplierByID, inputData FormUpdateSupplierInput) (Supplier, error) {
	supplier, err := s.repository.FindByID(input.ID)
	if err != nil {
		return supplier, err
	}

	supplier.Name = inputData.Name
	supplier.Address = inputData.Address
	supplier.Phone = inputData.Phone
	supplier.ContactPerson = inputData.ContactPerson
	updatedsupplier, err := s.repository.Update(supplier)
	if err != nil {
		return updatedsupplier, err
	}

	return updatedsupplier, nil
}

func (s *service) FindAll() ([]Supplier, error) {
	supplier, err := s.repository.FindAll()
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (s *service) GetSupplierDetail(input GetSupplierByID) (Supplier, error) {
	supplier, err := s.repository.FindByID(input.ID)
	if err != nil {
		return supplier, err
	}

	if supplier.ID == 0 {
		return supplier, errors.New("No supplier found")
	}

	return supplier, nil
}

func (s *service) SearchSUpplier(input GetSupplierSearchInput) ([]Supplier, error) {
	supplier, err := s.repository.SearchSupplier(input.Name)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}
func (s *service) DeleteSupplier(input GetSupplierByID) (Supplier, error) {
	supplier, err := s.repository.DeleteSupplier(input.ID)

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}
