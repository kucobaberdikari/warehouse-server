package product

type ProductFormatter struct {
	ID int `json:"id"`
	// SupplierID   int    `json:"supplier_id"`
	// SupplierName string `json:"supplier_name"`
	Supplier SupplierProductFormatter `json:"supplier"`
	// CategoryID   int    `json:"category_id"`
	// CategoryName string `json:"category_name"`
	Category CategoryProductFormatter ` json:"category"`
	SKU      string                   `json:"sku"`
	Name     string                   `json:"name"`
	Stock    int                      `json:"stock"`
	MinStock int                      `json:"min_stock"`
	Unit     string                   `json:"unit"`
	Image    string                   `json:"image"`
}

type SupplierProductFormatter struct {
	ID            int
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	// Phone         string `json:"phone"`
	// Address       string `json:"address"`
}

type CategoryProductFormatter struct {
	ID   int
	Name string `json:"name"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:       product.ID,
		SKU:      product.SKU,
		Name:     product.Name,
		Stock:    product.Stock,
		MinStock: product.MinStock,
		Unit:     product.Unit,
		// Supplier: product.Supplier,
		// Category: product.Category,
		Image: product.Image,
	}

	if product.Supplier.ID != 0 {
		formatter.Supplier = SupplierProductFormatter{
			ID:            product.Supplier.ID,
			Name:          product.Supplier.Name,
			ContactPerson: product.Supplier.ContactPerson,
			// Phone:         product.Supplier.Phone,
		}
	}

	if product.Category.ID != 0 {
		formatter.Category = CategoryProductFormatter{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		}
	}

	return formatter
}

type UpdateProductFormatter struct {
	ID int `json:"id"`
	// SupplierID   int    `json:"supplier_id"`
	// SupplierName string `json:"supplier_name"`
	// Supplier SupplierProductFormatter `json:"supplier"`
	// CategoryID   int    `json:"category_id"`
	// CategoryName string `json:"category_name"`
	// Category CategoryProductFormatter ` json:"category"`
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	MinStock int    `json:"min_stock"`
	Unit     string `json:"unit"`
	Image    string `json:"image"`
}

func FormatProductEdit(product Product) UpdateProductFormatter {
	formatter := UpdateProductFormatter{
		ID:       product.ID,
		SKU:      product.SKU,
		Name:     product.Name,
		Stock:    product.Stock,
		MinStock: product.MinStock,
		Image:    product.Image,
	}
	return formatter
}

type ProductUpdate struct {
	ID    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func FormatProducts(products []Product) []ProductFormatter {
	result := []ProductFormatter{}

	for _, product := range products {
		result = append(result, FormatProduct(product))
	}

	return result
}
