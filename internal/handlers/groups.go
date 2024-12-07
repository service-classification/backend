package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
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
//	@Param			group	body		services.NewGroup	true	"Group details"
//	@Success		201		{object}	models.Group
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/groups [post]
func (h *Handler) CreateGroup(c *gin.Context) {
	var group services.NewGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.GroupService.CreateGroup(group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// UpdateGroup godoc
//
//	@Summary		Update an existing group
//	@Description	Updates the details of an existing group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Group ID"
//	@Param			group	body		services.NewGroup	true	"Group details"
//	@Success		200		{object}	models.Group
//	@Failure		400		{object}	map[string]string	"Invalid input or group is used in services"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/groups/{id} [put]
func (h *Handler) UpdateGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	var group services.NewGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if exist any service with this group
	// if exist return error
	services, err := h.ServiceRepo.FindByGroupID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(services) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group is used in services"})
		return
	}

	model := &models.Group{
		ID:    group.ID,
		Title: group.Title,
	}
	if uint(groupID) != group.ID {
		err := h.GroupRepo.Delete(uint(groupID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = h.GroupRepo.Create(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := h.GroupRepo.Update(model); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//todo update group in the knowledge base

	c.JSON(http.StatusOK, model)
}

// DeleteGroup godoc
//
//	@Summary		Delete a group
//	@Description	Deletes a group by its ID. If the group is used in any services, it returns an error.
//	@Tags			Groups
//	@Param			id	path	int	true	"Group ID"
//	@Success		204	"Group deleted successfully"
//	@Failure		400	{object}	map[string]string	"Group is used in services"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/groups/{id} [delete]
func (h *Handler) DeleteGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	// check if exist any service with this group
	// if exist return error
	services, err := h.ServiceRepo.FindByGroupID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(services) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group is used in services"})
		return
	}

	if err := h.GroupRepo.Delete(uint(groupID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//todo remove group from the knowledge base

	c.JSON(http.StatusNoContent, nil)
}
