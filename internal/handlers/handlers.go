package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	ServiceRepo           repositories.ServiceRepository
	GroupRepo             repositories.GroupRepository
	ClassifiedServiceRepo repositories.ClassifiedServiceRepository
}

func NewHandler(serviceRepo repositories.ServiceRepository, groupRepo repositories.GroupRepository, csRepo repositories.ClassifiedServiceRepository) *Handler {
	return &Handler{
		ServiceRepo:           serviceRepo,
		GroupRepo:             groupRepo,
		ClassifiedServiceRepo: csRepo,
	}
}

// CreateService godoc
//
//	@Summary		Create a new service
//	@Description	Creates a new service with the provided details.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			service	body		models.Service	true	"Service details"
//	@Success		201		{object}	models.Service
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/services [post]
func (h *Handler) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ServiceRepo.Create(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, service)
}

// UpdateService godoc
//
//	@Summary		Update an existing service
//	@Description	Updates the details of an existing service by its ID.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Service ID"
//	@Param			service	body		models.Service	true	"Service details"
//	@Success		200		{object}	models.Service
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		404		{object}	map[string]string	"Service not found"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/services/{id} [put]
func (h *Handler) UpdateService(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.ID = uint(id)

	if err := h.ServiceRepo.Update(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

// DeleteService godoc
//
//	@Summary		Delete a service
//	@Description	Deletes a service by its ID.
//	@Tags			Services
//	@Param			id	path	int	true	"Service ID"
//	@Success		204
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/services/{id} [delete]
func (h *Handler) DeleteService(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.ServiceRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetServiceByID godoc
//
//	@Summary		Get a service by ID
//	@Description	Fetches the details of a service by its ID.
//	@Tags			Services
//	@Produce		json
//	@Param			id	path		int	true	"Service ID"
//	@Success		200	{object}	models.Service
//	@Failure		404	{object}	map[string]string	"Service not found"
//	@Router			/services/{id} [get]
func (h *Handler) GetServiceByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service, err := h.ServiceRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// ListServices godoc
//
//	@Summary		List all services
//	@Description	Fetches a list of services with pagination.
//	@Tags			Services
//	@Produce		json
//	@Param			offset	query		int	false	"Offset"	default(0)
//	@Param			limit	query		int	false	"Limit"		default(10)
//	@Success		200		{array}		models.Service
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/services [get]
func (h *Handler) ListServices(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	services, err := h.ServiceRepo.List(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}

// GetGroupByID godoc
//
//	@Summary		Get a group by ID
//	@Description	Retrieves a group by its ID.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Group ID"
//	@Success		200	{object}	models.Group
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/groups/{id} [get]
func (h *Handler) GetGroupByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := h.GroupRepo.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, group)
}

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
