package product

import "errors"

type Service interface {
	GetProduct(supplierID int) ([]Product, error)
	GetProducts() ([]Product, error)
	GetProductByID(input GetProductDetailInput) (Product, error)
	CreateProduct(input CreateProductInput) (Product, error)
	UpdateProduct(inputID GetProductDetailInput, inputData FormUpdateProductInput) (Product, error)
	DeleteProduct(inputID GetProductDetailInput) (Product, error)
	SearchProducts(input GetProductSearchInput) ([]Product, error)
	SearchSupplierProducts(input GetSupplierProductDetailInput) ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProduct(supplierID int) ([]Product, error) {
	products, err := s.repository.FindBySupplierID(supplierID)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetProducts() ([]Product, error) {
	product, err := s.repository.FindAll()
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) GetProductByID(input GetProductDetailInput) (Product, error) {
	product, err := s.repository.FindByID(input.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) CreateProduct(input CreateProductInput) (Product, error) {
	product := Product{}
	if input.MinStock >= input.Stock {
		return product, errors.New("Minimum stock must be lower than stock value!")
	}
	product.Name = input.Name
	product.SKU = input.SKU
	product.Stock = input.Stock
	product.MinStock = input.MinStock
	product.Unit = input.Unit
	product.SupplierID = input.SupplierID
	product.CategoryID = input.Category.ID

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) UpdateProduct(inputID GetProductDetailInput, inputData FormUpdateProductInput) (Product, error) {
	product, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return product, err
	}

	if product.Stock <= product.MinStock {
		return product, errors.New("product stock must be greater tahn minimum stock")
	}

	if product.Stock <= 0 || product.MinStock <= 0 {
		return product, errors.New("product must be greater")
	}

	product.Name = inputData.Name
	product.SKU = inputData.SKU
	product.Stock = inputData.Stock
	product.MinStock = inputData.MinStock
	product.Unit = inputData.Unit

	updateProduct, err := s.repository.Update(product)
	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (s *service) DeleteProduct(input GetProductDetailInput) (Product, error) {
	product, err := s.repository.DeleteProduct(input.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) SearchProducts(input GetProductSearchInput) ([]Product, error) {
	products, err := s.repository.SearchProduct(input.Name, input.SupplierID, input.CategoryID, input.SKU)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) SearchSupplierProducts(input GetSupplierProductDetailInput) ([]Product, error) {
	product, err := s.repository.FindBySupplierID(input.SupplierID)

	if err != nil {
		return product, err
	}

	return product, nil
}
