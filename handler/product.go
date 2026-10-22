package handler

import (
	"management-warehouse/helper"
	"management-warehouse/product"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service product.Service
	// authService    auth.Service
}

func NewProductHandler(service product.Service) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	search := c.Query("search")
	products, err := h.service.SearchProducts(product.GetProductSearchInput{
		Name:       search,
		SupplierID: search,
		CategoryID: search,
		SKU:        search,
	})
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	TotalRows := int64(len(products))

	PageStr := c.DefaultQuery("page", "1")
	LimitStr := c.DefaultQuery("limit", "10")

	page, errpage := strconv.Atoi(PageStr)
	if errpage != nil || page < 1 {
		page = 1
	}

	limit, errlimit := strconv.Atoi(LimitStr)
	if errlimit != nil || limit < 1 {
		limit = 6
	}
	offset := (page - 1) * limit
	TotalPage := int(math.Ceil(float64(TotalRows) / float64(limit)))

	pageProducts := make([]product.Product, 0)
	if offset < len(products) {
		end := offset + limit
		if end > len(products) {
			end = len(products)
		}
		pageProducts = products[offset:end]
	}

	Meta := gin.H{
		"total_rows": TotalRows,
		"total_page": TotalPage,
		"page":       page,
		"limit":      limit,
	}
	data := gin.H{
		"products": product.FormatProducts(pageProducts),
		"meta":     Meta,
	}
	response := helper.APIResponse("List of product", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	var input product.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.service.GetProductByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product detail", http.StatusOK, "success", product.FormatProduct(productDetail))
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetSupplierProducts(c *gin.Context) {

	var input product.GetSupplierProductDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.SearchSupplierProducts(input)
	if err != nil {
		response := helper.APIResponse("Error to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of product", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	var input product.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productdetail, err := h.service.GetProductByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product detail", http.StatusOK, "success", product.FormatProduct(productdetail))
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input product.CreateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.service.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create product", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var input product.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData product.FormUpdateProductInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateProduct, err := h.service.UpdateProduct(input, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update campaign", http.StatusOK, "success", product.FormatProductEdit(updateProduct))
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var input product.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.service.DeleteProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}
