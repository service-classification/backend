package handlers

import (
	"backend/internal/models"
	"log"
	"log/slog"
	"net/http"
	"sort"
	"strconv"
	"strings"
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

	invalidParameters, err := h.jenaService.ValidateService(c, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(invalidParameters) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters: " + strings.Join(invalidParameters, ", ")})
		return
	}

	if err := h.ServiceRepo.Create(service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go func() {
		payload := h.buildPayload(service.Parameters)
		predictions, err := callMLModel(payload)
		if err != nil {
			log.Println("Error calling ML model:", err)
			return
		}
		if len(predictions) > 0 {
			classID := uint(predictions[0].ClassID)
			correct, err := h.jenaService.ValidateClass(c, service, classID)
			if err != nil {
				log.Println("Error validating class:", err)
				return
			}
			if !correct {
				log.Println("Invalid class")
				return
			}

			class, err := h.ClassRepo.GetByID(classID)
			if err != nil {
				log.Println("Class not found:", err)
				return
			}

			service.Class = class
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

type assignClassRequest struct {
	ClassID *uint `json:"class_id,omitempty"`
}

// ApproveService godoc
//
//	@Summary		Approve a service
//	@Description	Approves a service by its ID. If a class ID is provided in the request body, it assigns the class to the service before approval.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Service ID"
//	@Param			class	body		assignClassRequest	true	"Class ID"
//	@Success		200		{object}	models.Service
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		404		{object}	map[string]string	"Service or class not found"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/services/{id}/approve [post]
func (h *Handler) ApproveService(c *gin.Context) {
	serviceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	req := assignClassRequest{}
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

	if req.ClassID == nil && service.Class == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class ID is required"})
		return
	}

	class := service.Class
	if req.ClassID != nil {
		class, err = h.ClassRepo.GetByID(*req.ClassID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
			return
		}
	}

	service.Class = class
	now := time.Now()
	service.ApprovedAt = &now

	if err := h.ServiceRepo.Update(service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.jenaService.AddService(c, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

// ListProposedClasses godoc
//
//	@Summary		List proposed classes for a service
//	@Description	Fetches a list of proposed classes for a service based on similar parameters.
//	@Tags			Services
//	@Produce		json
//	@Param			id	path		int	true	"Service ID"
//	@Success		200	{array}		models.Class
//	@Failure		400	{object}	map[string]string	"Invalid service ID"
//	@Failure		404	{object}	map[string]string	"Service not found"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/services/{id}/proposed_classes [get]
func (h *Handler) ListProposedClasses(c *gin.Context) {
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

	classes, err := h.jenaService.ProposedClasses(c, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := make([]proposedClassResponse, 0, len(classes))
	for _, class := range classes {
		resp := proposedClassResponse{}

		entityClass, err := h.ClassRepo.GetByID(class.ClassID)
		if err != nil {
			slog.Error("Category not found:", slog.Any("error", err))
			continue
		}
		resp.ClassID = entityClass.ID
		resp.Title = entityClass.Title

		resp.SimilarParameters = class.MatchingParameterNums
		resp.SimilarServices = len(class.SimilarServices)

		result = append(result, resp)
	}

	// sort result first by similar services and then by similar parameters
	sort.Slice(result, func(i, j int) bool {
		if result[i].SimilarServices == result[j].SimilarServices {
			return result[i].SimilarParameters > result[j].SimilarParameters
		}
		return result[i].SimilarServices > result[j].SimilarServices
	})

	c.JSON(http.StatusOK, result)
}

type proposedClassResponse struct {
	ClassID           uint   `json:"class_id"`
	Title             string `json:"title"`
	SimilarParameters int    `json:"similar_parameters"`
	SimilarServices   int    `json:"similar_services"`
}
