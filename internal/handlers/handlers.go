package handlers

import (
	"backend/internal/apache_jena"
	"backend/internal/repositories"
	"backend/internal/services"
)

type Handler struct {
	ServiceRepo   repositories.ServiceRepository
	ClassRepo     repositories.ClassRepository
	ParameterRepo repositories.ParameterRepository

	ClassService     *services.ClassService
	ParameterService *services.ParameterService
	jenaService      *apache_jena.Service
}

func NewHandler(serviceRepo repositories.ServiceRepository, classRepository repositories.ClassRepository, paramRepo repositories.ParameterRepository, service *apache_jena.Service, parameterService *services.ParameterService, classService *services.ClassService) *Handler {
	return &Handler{
		ServiceRepo:      serviceRepo,
		ClassRepo:        classRepository,
		ParameterRepo:    paramRepo,
		ClassService:     classService,
		ParameterService: parameterService,
		jenaService:      service,
	}
}
