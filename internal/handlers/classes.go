package handlers

import (
	"backend/internal/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListClasses godoc
//
//	@Summary		List classes with pagination
//	@Description	Retrieves a list of classes with pagination.
//	@Tags			Classes
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int	false	"Offset"	default(0)
//	@Param			limit	query		int	false	"Limit"		default(10)
//	@Success		200		{array}		models.Class
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/classes [get]
func (h *Handler) ListClasses(c *gin.Context) {
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

	classes, err := h.ClassRepo.List(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}

// CreateClass godoc
//
//	@Summary		Create a new class
//	@Description	Creates a new class with the provided details.
//	@Tags			Classes
//	@Accept			json
//	@Produce		json
//	@Param			class	body		models.ClassView	true	"Class details"
//	@Success		201		{object}	models.Class
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/classes [post]
func (h *Handler) CreateClass(c *gin.Context) {
	var class models.ClassView
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.ClassService.CreateClass(class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.jenaService.AddClass(c, class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// GetClassByID godoc
//
//	@Summary		Get a class by ID
//	@Description	Retrieves a class by its ID.
//	@Tags			Classes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Class ID"
//	@Success		200	{object}	models.ClassView
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/classes/{id} [get]
func (h *Handler) GetClassByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	class, err := h.ClassRepo.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	constraints, err := h.jenaService.GetClassConstraints(c, class.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	view := models.ClassView{
		ID:                class.ID,
		Title:             class.Title,
		AllowedParameters: constraints,
	}

	c.JSON(http.StatusOK, view)
}

// UpdateClass godoc
//
//	@Summary		Update an existing class
//	@Description	Updates the details of an existing class.
//	@Tags			Classes
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Class ID"
//	@Param			class	body		models.ClassView	true	"Class details"
//	@Success		200		{object}	models.Class
//	@Failure		400		{object}	map[string]string	"Invalid input or class is used in services"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/classes/{id} [put]
func (h *Handler) UpdateClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	var class models.ClassView
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if exist any service with this class
	// if exist return error
	services, err := h.ServiceRepo.FindByClassID(uint(classID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(services) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class is used in services"})
		return
	}

	model := &models.Class{
		ID:    class.ID,
		Title: class.Title,
	}
	if uint(classID) != class.ID {
		err := h.ClassRepo.Delete(uint(classID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = h.ClassRepo.Create(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := h.ClassRepo.Update(model); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = h.jenaService.UpdateClass(c, class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model)
}

// DeleteClass godoc
//
//	@Summary		Delete a class
//	@Description	Deletes a class by its ID. If the class is used in any services, it returns an error.
//	@Tags			Classes
//	@Param			id	path	int	true	"Class ID"
//	@Success		204	"Class deleted successfully"
//	@Failure		400	{object}	map[string]string	"Class is used in services"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/classes/{id} [delete]
func (h *Handler) DeleteClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	// check if exist any service with this class
	// if exist return error
	services, err := h.ServiceRepo.FindByClassID(uint(classID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(services) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class is used in services"})
		return
	}

	if err := h.ClassRepo.Delete(uint(classID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.jenaService.DeleteClass(c, uint(classID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
