package services

import (
	"backend/internal/apache_jena"
	"backend/internal/models"
	"backend/internal/repositories"
	"context"
)

type ParameterService struct {
	ParameterRepository repositories.ParameterRepository
	jenaService         *apache_jena.Service
}

func NewParameterService(parameterRepo repositories.ParameterRepository, service *apache_jena.Service) *ParameterService {
	return &ParameterService{
		ParameterRepository: parameterRepo,
		jenaService:         service,
	}
}

func (s *ParameterService) CreateParameter(parameter models.ParameterView) (*models.Parameter, error) {
	model := &models.Parameter{
		ID:    parameter.ID,
		Title: parameter.Title,
	}
	err := s.ParameterRepository.Create(model)
	if err != nil {
		return model, err
	}

	err = s.jenaService.AddParameter(context.TODO(), parameter)
	if err != nil {
		return model, err
	}

	return model, nil
}
