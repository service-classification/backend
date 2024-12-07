package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListParameters godoc
//
//	@Summary		List parameters with pagination
//	@Description	Retrieves a list of parameters with pagination.
//	@Tags			Parameters
//	@Produce		json
//	@Param			offset	query		int	false	"Offset"	default(0)
//	@Param			limit	query		int	false	"Limit"		default(10)
//	@Success		200		{array}		models.Parameter
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/parameters [get]
func (h *Handler) ListParameters(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	parameters, err := h.ParameterRepo.List(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parameters)
}

// CreateParameter godoc
//
//	@Summary		Create a new parameter
//	@Description	Creates a new parameter with the provided details.
//	@Tags			Parameters
//	@Accept			json
//	@Produce		json
//	@Param			parameter	body		services.NewParameter	true	"Parameter details"
//	@Success		201			{object}	models.Parameter
//	@Failure		400			{object}	map[string]string	"Invalid input"
//	@Failure		500			{object}	map[string]string	"Internal server error"
//	@Router			/parameters [post]
func (h *Handler) CreateParameter(c *gin.Context) {
	var parameter services.NewParameter
	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.ParameterService.CreateParameter(parameter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// UpdateParameter godoc
//
//	@Summary		Update an existing parameter
//	@Description	Updates an existing parameter with the provided details.
//	@Tags			Parameters
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string					true	"Parameter ID"
//	@Param			parameter	body		services.NewParameter	true	"Parameter details"
//	@Success		200			{object}	models.Parameter
//	@Failure		400			{object}	map[string]string	"Invalid input or parameter is used in services"
//	@Failure		500			{object}	map[string]string	"Internal server error"
//	@Router			/parameters/{id} [put]
func (h *Handler) UpdateParameter(c *gin.Context) {
	parameterID := c.Param("id")

	var parameter services.NewParameter
	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if exist any service with this parameter
	// if exist return error
	services, err := h.ServiceRepo.FindByParameterID(parameterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(services) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter is used in services"})
		return
	}

	model := &models.Parameter{
		ID:    parameter.ID,
		Title: parameter.Title,
	}

	if parameter.ID != parameterID {
		if err := h.ParameterRepo.Delete(parameterID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := h.ParameterRepo.Create(model); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := h.ParameterRepo.Update(model); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//todo update parameter in the knowledge base

	c.JSON(http.StatusOK, model)
}

// DeleteParameter godoc
//
//	@Summary		Delete a parameter
//	@Description	Deletes a parameter by its ID. If the parameter is used in any services, it returns an error.
//	@Tags			Parameters
//	@Param			id	path	string	true	"Parameter ID"
//	@Success		204	"Parameter deleted successfully"
//	@Failure		400	{object}	map[string]string	"Parameter is used in services"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/parameters/{id} [delete]
func (h *Handler) DeleteParameter(c *gin.Context) {
	parameterID := c.Param("id")

	// check if exist any service with this parameter
	// if exist return error
	services, err := h.ServiceRepo.FindByParameterID(parameterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(services) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter is used in services"})
		return
	}

	if err := h.ParameterRepo.Delete(parameterID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//todo delete parameter from the knowledge base

	c.JSON(http.StatusNoContent, nil)
}
