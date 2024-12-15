package services

import (
	"backend/internal/apache_jena"
	"backend/internal/models"
	"backend/internal/repositories"
	"context"
)

type ClassService struct {
	ClassRepository repositories.ClassRepository
	jenaService     *apache_jena.Service
}

func NewClassService(classRepository repositories.ClassRepository, service *apache_jena.Service) *ClassService {
	return &ClassService{
		ClassRepository: classRepository,
		jenaService:     service,
	}
}

func (s *ClassService) CreateClass(classView models.ClassView, new bool) (*models.Class, error) {
	model := &models.Class{
		ID:    classView.ID,
		Title: classView.Title,
		New:   new,
	}
	err := s.ClassRepository.Create(model)
	if err != nil {
		return model, err
	}

	err = s.jenaService.AddClass(context.TODO(), classView)
	if err != nil {
		return model, err
	}

	return model, nil
}
