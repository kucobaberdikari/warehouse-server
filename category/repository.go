package category

import "gorm.io/gorm"

type Repository interface {
	Save(category Category) (Category, error)
	FindByID(ID int) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	DeleteCategory(ID int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByID(ID int) (Category, error) {
	var category Category

	err := r.db.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindAll() ([]Category, error) {
	var category []Category

	err := r.db.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
func (r *repository) Update(category Category) (Category, error) {

	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) DeleteCategory(ID int) (Category, error) {
	var category Category
	err := r.db.Where("id = ?", ID).Find(&category).Delete(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
