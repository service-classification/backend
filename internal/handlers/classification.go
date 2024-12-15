package handlers

import (
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Helper functions
func (h *Handler) buildPayload(parameters []models.Parameter) (map[string]int, error) {
	supportedParams, err := h.ParameterRepo.ListSupportedParameters()
	if err != nil {
		return nil, err
	}

	payload := make(map[string]int, len(supportedParams))
	for _, param := range supportedParams {
		payload[param] = 0
	}

	for _, param := range parameters {
		if _, ok := payload[param.ID]; ok {
			payload[param.ID] = 1
		}
	}
	return payload, nil
}

func (h *Handler) callMLModel(parameters []models.Parameter) ([]Prediction, error) {
	payload, err := h.buildPayload(parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to build payload: %w", err)
	}

	url := os.Getenv("ML_MODEL_URL")
	token := os.Getenv("BEARER_TOKEN")

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

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
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
	ClassID     int     `json:"group_id"`
	Probability float64 `json:"probability"`
}
