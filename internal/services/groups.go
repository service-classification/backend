package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type ClassService struct {
	ClassRepository repositories.ClassRepository
}

func NewClassService(classRepository repositories.ClassRepository) *ClassService {
	return &ClassService{
		ClassRepository: classRepository,
	}
}

type ClassView struct {
	ID                uint     `json:"id" example:"3042"`
	Title             string   `json:"title"`
	AllowedParameters []string `json:"allowed_parameters" example:"mob_inet,fix_ctv,voice_fix"`
}

func (s *ClassService) CreateClass(classView ClassView) (*models.Class, error) {
	model := &models.Class{
		ID:    classView.ID,
		Title: classView.Title,
	}
	err := s.ClassRepository.Create(model)
	if err != nil {
		return model, err
	}

	//todo add new class to the knowledge base

	return model, nil
}
