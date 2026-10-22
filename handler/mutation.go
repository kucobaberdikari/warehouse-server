package handler

import (
	"management-warehouse/helper"
	"management-warehouse/mutation"
	"management-warehouse/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MutationHandler struct {
	service mutation.Service
}

func NewMutationHandler(service mutation.Service) *MutationHandler {
	return &MutationHandler{service}
}

func (h *MutationHandler) CreateMutation(c *gin.Context) {
	var input mutation.CreateMutationInput

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create mutation", http.StatusUnprocessableEntity, "error", errorMessage)
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
	input.User = user

	newMutation, err := h.service.CreateMutation(input)
	if err != nil {
		response := helper.APIResponse("Error to get mutation", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of mutation", http.StatusOK, "success", mutation.FormatMutation(newMutation))
	c.JSON(http.StatusOK, response)
}

func (h *MutationHandler) GetMutations(c *gin.Context) {
	mutations, err := h.service.GetMutation()

	if err != nil {
		response := helper.APIResponse("Error to get mutation", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of mutation", http.StatusOK, "success", mutation.FormatMutations(mutations))
	c.JSON(http.StatusOK, response)
}

func (h *MutationHandler) GetMutationByID(c *gin.Context) {
	var input mutation.GetMutationDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get mutation", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mutationDetail, err := h.service.GetMutationByID(input)
	if err != nil {
		response := helper.APIResponse("Error to get mutation", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of mutation", http.StatusOK, "success", mutation.FormatMutation(mutationDetail))
	c.JSON(http.StatusOK, response)

}
