package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type ParameterService struct {
	ParameterRepository repositories.ParameterRepository
}

func NewParameterService(parameterRepo repositories.ParameterRepository) *ParameterService {
	return &ParameterService{
		ParameterRepository: parameterRepo,
	}
}

type ParameterView struct {
	ID             string `json:"id" example:"fix_ctv"`
	Title          string `json:"title"`
	AllowedClasses []int  `json:"allowed_classes" example:"1,1033,3023"`
}

func (s *ParameterService) CreateParameter(parameter ParameterView) (*models.Parameter, error) {
	model := &models.Parameter{
		ID:    parameter.ID,
		Title: parameter.Title,
	}
	err := s.ParameterRepository.Create(model)
	if err != nil {
		return model, err
	}

	//todo add new parameter to the knowledge base

	return model, nil
}
