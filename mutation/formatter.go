package mutation

import (
	"time"
)

type MutationFormatter struct {
	ID        int              `json:"id"`
	Product   ProductFormatter `json:"product"`
	Type      string           `json:"type"`
	Quantity  int              `json:"quantity"`
	Note      string           `json:"note"`
	CreatedAt time.Time        `json:"created_at"`
	User      UserFormatter    `json:"user"`
}
type UserFormatter struct {
	// ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductFormatter struct {
	ID        int                      `json:"id"`
	Supplier  SupplierProductFormatter `json:"supplier"`
	Category  CategoryProductFormatter `json:"category"`
	SKU       string                   `json:"sku"`
	Name      string                   `json:"name"`
	Stock     int                      `json:"stock"`
	MinStock  int                      `json:"min_stock"`
	Unit      string                   `json:"unit"`
	Image     string                   `json:"image"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type SupplierProductFormatter struct {
	// ID            int
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	// Phone         string `json:"phone"`
	// Address       string `json:"address"`
}

type CategoryProductFormatter struct {
	// ID   int
	Name string `json:"name"`
}

func FormatMutation(mutation Mutation) MutationFormatter {
	formatter := MutationFormatter{
		ID:        mutation.ID,
		Type:      string(mutation.Type),
		Quantity:  mutation.Quantity,
		Note:      mutation.Note,
		CreatedAt: mutation.CreatedAt,
	}
	if mutation.User.ID != 0 {
		formatter.User = UserFormatter{
			// ID:   mutation.User.ID,
			Name: mutation.User.Name,
		}
	}

	if mutation.Product.ID != 0 {
		formatter.Product = ProductFormatter{
			ID: mutation.Product.ID,
			Supplier: SupplierProductFormatter{
				// ID:            mutation.Product.Supplier.ID,
				Name:          mutation.Product.Supplier.Name,
				ContactPerson: mutation.Product.Supplier.ContactPerson,
			},
			Category: CategoryProductFormatter{
				// ID:   mutation.Product.Category.ID,
				Name: mutation.Product.Category.Name,
			},
			SKU:       mutation.Product.SKU,
			Name:      mutation.Product.Name,
			Stock:     mutation.Product.Stock,
			MinStock:  mutation.Product.MinStock,
			Unit:      mutation.Product.Unit,
			Image:     mutation.Product.Image,
			UpdatedAt: mutation.Product.UpdatedAt,
		}
	}

	return formatter
}

func FormatMutations(mutation []Mutation) []MutationFormatter {
	result := []MutationFormatter{}

	for _, mutation := range mutation {
		result = append(result, FormatMutation(mutation))
	}

	return result
}
