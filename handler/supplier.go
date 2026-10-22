package handler

import (
	"management-warehouse/auth"
	"management-warehouse/helper"
	"management-warehouse/supplier"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type supplierHandler struct {
	supplierService supplier.Service
	authService     auth.Service
}

func NewSupplierHandler(supplierService supplier.Service, authService auth.Service) *supplierHandler {
	return &supplierHandler{supplierService, authService}
}

func (h *supplierHandler) RegisterSupplier(c *gin.Context) {
	var input supplier.RegisterSupplierInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSupplier, err := h.supplierService.RegisterSupplier(input)

	if err != nil {
		response := helper.APIResponse("Register supplier failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := supplier.FormatSupplier(newSupplier)

	response := helper.APIResponse("Supplier has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) GetAllSupplier(c *gin.Context) {
	search := c.Query("search")
	suppliers, err := h.supplierService.SearchSUpplier(supplier.GetSupplierSearchInput{
		Name: search,
	})
	if err != nil {
		response := helper.APIResponse("Error to get supplier", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	TotalRows := int64(len(suppliers))

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
	pagesupplier := make([]supplier.Supplier, 0)

	if offset < len(suppliers) {
		end := offset + limit
		if end > len(suppliers) {
			end = len(suppliers)
		}
		pagesupplier = suppliers[offset:end]
	}
	Meta := gin.H{
		"total_rows": TotalRows,
		"total_page": TotalPage,
		"page":       page,
		"limit":      limit,
	}
	data := gin.H{
		"supplier": supplier.FormatSuppliers(pagesupplier),
		"meta":     Meta,
	}
	response := helper.APIResponse("List of supplier", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
func (h *supplierHandler) GetSupplierByID(c *gin.Context) {
	var input supplier.GetSupplierByID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	supplierdetail, err := h.supplierService.GetSupplierByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Supplier detail", http.StatusOK, "success", supplier.FormatSupplier(supplierdetail))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) UpdateSupplier(c *gin.Context) {
	var input supplier.GetSupplierByID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData supplier.FormUpdateSupplierInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateSupplier, err := h.supplierService.Updatesupplier(input, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update supplier", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Supplier detail", http.StatusOK, "success", supplier.FormatSupplier(updateSupplier))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) DeleteSupplier(c *gin.Context) {
	var input supplier.GetSupplierByID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	supplier, err := h.supplierService.DeleteSupplier(input)
	if err != nil {
		response := helper.APIResponse("Failed to update supplier", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Supplier detail", http.StatusOK, "success", supplier)
	c.JSON(http.StatusOK, response)
}
