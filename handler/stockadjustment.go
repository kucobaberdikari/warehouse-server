package handler

import (
	"management-warehouse/helper"
	"management-warehouse/stockadjustment"
	"management-warehouse/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockAdjustmentHandler struct {
	service stockadjustment.Service
}

func NewStockAdjustmentHandler(service stockadjustment.Service) *StockAdjustmentHandler {
	return &StockAdjustmentHandler{service}
}
func (h *StockAdjustmentHandler) Create(c *gin.Context) {
	var input stockadjustment.CreateStockAdjustmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create stock adjusment", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser, ok := c.Get("currentUser")
	if !ok {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	user, ok := currentUser.(user.User)
	if !ok || user.ID <= 0 {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	input.CreatedBy = user.ID

	adjustment, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create stock adjusment", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success to create product", http.StatusOK, "success", stockadjustment.StockAdjustmentFormat(adjustment))
	c.JSON(http.StatusOK, response)
}

func (h *StockAdjustmentHandler) GetAll(c *gin.Context) {

	adjustments, err := h.service.GetAll()
	if err != nil {
		response := helper.APIResponse("Failed to get list of stock adjustment", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get stock adjustment", http.StatusOK, "success", stockadjustment.StockAdjustmentsFormat(adjustments))
	c.JSON(http.StatusOK, response)
}

func (h *StockAdjustmentHandler) GetByID(c *gin.Context) {
	var input stockadjustment.GetStockAdjustmentInput

	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.APIResponse("Failed to get detail of stock adjustment", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	adjustment, err := h.service.GetByID(input)
	if err != nil {
		response := helper.APIResponse("stock adjustment not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get stock adjustment detail", http.StatusOK, "success", stockadjustment.StockAdjustmentFormat(adjustment))
	c.JSON(http.StatusOK, response)
}

func (h *StockAdjustmentHandler) GetByProductID(c *gin.Context) {

	// var input stockadjustment.GetStockAdjustmentInput
	// productID := c.Param("product_id")

	// parse int
	// call service.GetByProductID(productID)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   nil,
	})
}
