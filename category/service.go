package category

import "time"

type Service interface {
	CreateCategory(input FormCategoryInput) (Category, error)
	GetCategoryByID(input GetCategoryInput) (Category, error)
	UpdateCategory(inputID GetCategoryInput, inputData FormCategoryInput) (Category, error)
	GetAllCategory() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCategory(input FormCategoryInput) (Category, error) {
	category := Category{}
	category.Name = input.Name
	category.CreatedAt = time.Now()

	newCat, err := s.repository.Save(category)
	if err != nil {
		return newCat, err
	}
	return newCat, nil
}

func (s *service) GetCategoryByID(input GetCategoryInput) (Category, error) {
	category, err := s.repository.FindByID(input.ID)
	if err != nil {
		return category, err
	}
	return category, nil
}
func (s *service) UpdateCategory(inputID GetCategoryInput, inputData FormCategoryInput) (Category, error) {
	category, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return category, err
	}

	category.Name = inputData.Name

	updatecat, err := s.repository.Update(category)
	if err != nil {
		return updatecat, err
	}

	return updatecat, nil
}

func (s *service) GetAllCategory() ([]Category, error) {
	category, err := s.repository.FindAll()
	if err != nil {
		return category, err
	}
	return category, nil
}
