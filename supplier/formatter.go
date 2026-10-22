package supplier

type SupplierFormatter struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
}

func FormatSupplier(supplier Supplier) SupplierFormatter {
	formatter := SupplierFormatter{
		ID:            supplier.ID,
		Name:          supplier.Name,
		ContactPerson: supplier.ContactPerson,
		Phone:         supplier.Phone,
		Address:       supplier.Address,
	}
	return formatter
}

func FormatSupplierName(supplier Supplier) SupplierFormatter {
	formatter := SupplierFormatter{
		ID:    supplier.ID,
		Name:  supplier.Name,
		Phone: supplier.Phone,
	}
	return formatter
}

func FormatSuppliers(supplier []Supplier) []SupplierFormatter {
	result := []SupplierFormatter{}
	for _, supplier := range supplier {
		result = append(result, FormatSupplier(supplier))
	}
	return result
}
