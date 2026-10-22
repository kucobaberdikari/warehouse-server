package supplier

type RegisterSupplierInput struct {
	Name          string `json:"name" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Address       string `json:"address" binding:"required"`
}

type FormUpdateSupplierInput struct {
	// ID            int
	Name          string `json:"name" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Address       string `json:"address" binding:"required"`
}

type GetSupplierByID struct {
	ID int `uri:"id" binding:"required"`
}

type GetSupplierSearchInput struct {
	Name string `form:"name"`
}
