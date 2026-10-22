package category

type FormCategoryInput struct {
	// ID   int    `uri:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	// CreatedAt time.Time `json:"created_at" binding:"required"`
}

type GetCategoryInput struct {
	ID int `uri:"id" binding:"required"`
}

type FormCategoryUpdate struct {
	ID   int
	Name string `json:"name" binding:"required"`
	// CreatedAt time.Time `json:"created_at" binding:"required"`
}
