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
	// from list filter only allowed parameters:
	payload := map[string]int{
		"mob_inet":                0,
		"fix_inet":                0,
		"fix_ctv":                 0,
		"fix_ictv":                0,
		"voice_mob":               0,
		"voice_fix":               0,
		"sms":                     0,
		"csd":                     0,
		"iot":                     0,
		"mms":                     0,
		"roaming":                 0,
		"mg":                      0,
		"mn":                      0,
		"mts":                     0,
		"conc":                    0,
		"fix_op":                  0,
		"vsr_roam":                0,
		"national_roam":           0,
		"mn_roam":                 0,
		"voice_ap":                0,
		"voice_fee":               0,
		"period_service":          0,
		"one_time_service":        0,
		"dop_service":             0,
		"content":                 0,
		"services_service":        0,
		"keo_sale":                0,
		"discount":                0,
		"only_inbound":            0,
		"sms_a2p":                 0,
		"sms_gross":               0,
		"skoring":                 0,
		"other_service":           0,
		"voice_mail":              0,
		"geo":                     0,
		"ep_for_number":           0,
		"ep_for_line":             0,
		"one_time_fee_for_number": 0,
		"equipment rent":          0,
		"add_package":             0,
	}

	for _, param := range parameters {
		if _, ok := payload[param.ID]; ok {
			payload[param.ID] = 1
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

func callMLModel(payload map[string]int) ([]Prediction, error) {
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
		Predictions []Prediction `json:"predictions"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Predictions, nil
}

type Prediction struct {
	GroupID     int     `json:"group_id"`
	Probability float64 `json:"probability"`
}
