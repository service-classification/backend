package handlers

import (
	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListGroups godoc
//
//	@Summary		List groups with pagination
//	@Description	Retrieves a list of groups with pagination.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int	false	"Offset"	default(0)
//	@Param			limit	query		int	false	"Limit"		default(10)
//	@Success		200		{array}		models.Group
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/groups [get]
func (h *Handler) ListGroups(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	groups, err := h.GroupRepo.List(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// CreateGroup godoc
//
//	@Summary		Create a new group
//	@Description	Creates a new group with the provided details.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			group	body		models.Group	true	"Group details"
//	@Success		201		{object}	models.Group
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/groups [post]
func (h *Handler) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.GroupRepo.Create(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//todo add new group to the knowledge base

	c.JSON(http.StatusCreated, group)
}
