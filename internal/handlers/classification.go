package handlers

import (
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ClassifyService godoc
//
//	@Summary		Classify a service
//	@Description	Classifies a service by its ID using the ML model.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			service	query		int						true	"Service ID"
//	@Success		200		{object}	map[string]interface{}	"Predictions"
//	@Failure		404		{object}	map[string]string		"Service not found"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/classify [post]
func (h *Handler) ClassifyService(c *gin.Context) {
	serviceID, _ := strconv.Atoi(c.Query("service"))
	service, err := h.ServiceRepo.GetByID(uint(serviceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	// Build the payload for the ML model
	payload := h.buildPayload(service.Parameters)

	// Call the ML model API
	predictions, err := callMLModel(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"predictions": predictions})
}

// AssignGroupToService godoc
//
//	@Summary		Assign a group to a service
//	@Description	Assigns a group to a service by their IDs.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Service ID"
//	@Param			group	body		assignGroupRequest	true	"Group ID"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/services/{id}/group [post]
func (h *Handler) AssignGroupToService(c *gin.Context) {
	// Parse service ID from URL parameter
	serviceIDParam := c.Param("id")
	serviceIDUint64, err := strconv.ParseUint(serviceIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}
	serviceID := uint(serviceIDUint64)

	// Parse group ID from request body
	req := assignGroupRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}
	groupID := req.GroupID

	// Validate that the service exists
	_, err = h.ServiceRepo.GetByID(serviceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	// Validate that the group exists
	_, err = h.GroupRepo.GetByID(groupID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	// Assign group to service
	err = h.ClassifiedServiceRepo.AssignGroupToService(serviceID, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign group to service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group assigned to service successfully"})
}

type assignGroupRequest struct {
	GroupID uint `json:"group_id" example:"2"`
}

// Helper functions
func (h *Handler) buildPayload(parameters []models.Parameter) map[string]int {
	list, err := h.ParameterRepo.List(0, 100000)
	if err != nil {
		return map[string]int{}
	}

	payload := make(map[string]int)
	for _, param := range list {
		if contains(parameters, param) {
			payload[param.Code] = 1
		} else {
			payload[param.Code] = 0
		}
	}
	return payload
}

func contains(slice []models.Parameter, item models.Parameter) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func callMLModel(payload map[string]int) ([]map[string]interface{}, error) {
	url := os.Getenv("ML_MODEL_URL")   // todo from config
	token := os.Getenv("BEARER_TOKEN") // todo from config

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(bodyBytes))
	}

	var result struct {
		Predictions []map[string]interface{} `json:"predictions"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Predictions, nil
}
