package handler

import (
	"management-warehouse/category"
	"management-warehouse/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service category.Service
}

func NewCategoryHandler(service category.Service) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input category.FormCategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create new category failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newcategory, err := h.service.CreateCategory(input)
	if err != nil {
		response := helper.APIResponse("Failed to create category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create product", http.StatusOK, "success", category.FormCategory(newcategory))
	c.JSON(http.StatusOK, response)

}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	var input category.GetCategoryInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoryDetail, err := h.service.GetCategoryByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("category", http.StatusOK, "success", category.FormCategory(categoryDetail))
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.GetAllCategory()
	if err != nil {
		response := helper.APIResponse("Failed to get category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("category", http.StatusOK, "success", category.FormatCategories(categories))
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	var input category.GetCategoryInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoryDetail, err := h.service.GetCategoryByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("category detail", http.StatusOK, "success", category.FormatCategory(categoryDetail))
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var input category.GetCategoryInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData category.FormCategoryInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatecat, err := h.service.UpdateCategory(input, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success to update category", http.StatusOK, "success", category.FormatCategory(updatecat))
	c.JSON(http.StatusOK, response)
}
