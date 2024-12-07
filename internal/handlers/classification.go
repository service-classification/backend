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
