package handlers

import (
	"backend/internal/repositories"
	"backend/internal/services"
)

type Handler struct {
	ServiceRepo   repositories.ServiceRepository
	ClassRepo     repositories.ClassRepository
	ParameterRepo repositories.ParameterRepository

	ClassService     *services.ClassService
	ParameterService *services.ParameterService
}

func NewHandler(serviceRepo repositories.ServiceRepository, classRepository repositories.ClassRepository, paramRepo repositories.ParameterRepository) *Handler {
	classService := services.NewClassService(classRepository)
	parameterService := services.NewParameterService(paramRepo)
	return &Handler{
		ServiceRepo:      serviceRepo,
		ClassRepo:        classRepository,
		ParameterRepo:    paramRepo,
		ClassService:     classService,
		ParameterService: parameterService,
	}
}
