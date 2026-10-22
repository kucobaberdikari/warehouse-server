package category

import "time"

type CategoryFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{}
	formatter.ID = category.ID
	formatter.Name = category.Name
	// formatter.CreatedAt = category.CreatedAt
	return formatter
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:   category.ID,
		Name: category.Name,
	}
	return formatter
}

func FormatCategories(category []Category) []CategoryFormatter {
	categoriesformatter := []CategoryFormatter{}

	for _, categories := range category {
		categoryFormatter := FormCategory(categories)
		categoriesformatter = append(categoriesformatter, categoryFormatter)
	}
	return categoriesformatter
}
