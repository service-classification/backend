package handlers

import (
	"backend/internal/models"
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
//	@Param			parameter	body		models.Parameter	true	"Parameter details"
//	@Success		201			{object}	models.Parameter
//	@Failure		400			{object}	map[string]string	"Invalid input"
//	@Failure		500			{object}	map[string]string	"Internal server error"
//	@Router			/parameters [post]
func (h *Handler) CreateParameter(c *gin.Context) {
	var parameter models.Parameter
	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ParameterRepo.Create(&parameter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//todo add new parameter to the knowledge base

	c.JSON(http.StatusCreated, parameter)
}
