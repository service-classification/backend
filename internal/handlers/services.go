package handlers

import (
	"backend/internal/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type NewService struct {
	Title      string   `json:"title" binding:"required"`
	Parameters []string `json:"parameters" binding:"required,dive,required"`
}

// CreateService godoc
//
//	@Summary		Create a new service
//	@Description	Creates a new service with the provided details.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			service	body		NewService	true	"Service details"
//	@Success		201		{object}	models.Service
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/services [post]
func (h *Handler) CreateService(c *gin.Context) {
	var newService NewService
	if err := c.ShouldBindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := make([]models.Parameter, 0, len(newService.Parameters))
	for _, param := range newService.Parameters {
		parameter, err := h.ParameterRepo.GetByID(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
			return
		}
		params = append(params, *parameter)
	}

	service := &models.Service{
		Title:      newService.Title,
		Parameters: params,
	}

	if err := h.ServiceRepo.Create(service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//todo classify the service

	go func() {
		payload := h.buildPayload(service.Parameters)
		predictions, err := callMLModel(payload)
		if err != nil {
			log.Println("Error calling ML model:", err)
			return
		}
		if len(predictions) > 0 {
			groupID, err := strconv.Atoi(predictions[0].GroupID)
			if err != nil {
				log.Println("Invalid group ID:", err)
				return
			}
			group, err := h.GroupRepo.GetByID(uint(groupID))
			if err != nil {
				log.Println("Group not found:", err)
				return
			}
			service.Group = group
			err = h.ServiceRepo.Update(service)
			if err != nil {
				log.Println("Error updating service:", err)
				return
			}
		}
	}()

	c.JSON(http.StatusCreated, service)
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

type assignGroupRequest struct {
	GroupID *uint `json:"group_id,omitempty"`
}

// ApproveService godoc
//
//	@Summary		Approve a service
//	@Description	Approves a service by its ID. If a group ID is provided in the request body, it assigns the group to the service before approval.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Service ID"
//	@Param			group	body		assignGroupRequest	true	"Group ID"
//	@Success		200		{object}	models.Service
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		404		{object}	map[string]string	"Service or group not found"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/services/{id}/approve [post]
func (h *Handler) ApproveService(c *gin.Context) {
	serviceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	req := assignGroupRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	service, err := h.ServiceRepo.GetByID(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	if service.ApprovedAt != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Service is already approved"})
		return
	}

	if req.GroupID == nil && service.Group == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is required"})
		return
	}

	group := service.Group
	if req.GroupID != nil {
		group, err = h.GroupRepo.GetByID(*req.GroupID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		}
	}

	service.Group = group
	now := time.Now()
	service.ApprovedAt = &now

	if err := h.ServiceRepo.Update(service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//todo add approved service to the knowledge base

	c.JSON(http.StatusOK, service)
}

// ListProposedGroups godoc
//
//	@Summary		List proposed groups for a service
//	@Description	Fetches a list of proposed groups for a service based on similar parameters.
//	@Tags			Services
//	@Produce		json
//	@Param			id	path		int	true	"Service ID"
//	@Success		200	{array}		models.Group
//	@Failure		400	{object}	map[string]string	"Invalid service ID"
//	@Failure		404	{object}	map[string]string	"Service not found"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/services/{id}/proposed_groups [get]
func (h *Handler) ListProposedGroups(c *gin.Context) {
	serviceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	service, err := h.ServiceRepo.GetByID(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	//todo find approved services with similar parameters
	_ = service

	groups, err := h.GroupRepo.List(0, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}
