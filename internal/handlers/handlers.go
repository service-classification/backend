package handlers

import (
	"backend/internal/repositories"
	"backend/internal/services"
)

type Handler struct {
	ServiceRepo   repositories.ServiceRepository
	GroupRepo     repositories.GroupRepository
	ParameterRepo repositories.ParameterRepository

	GroupService     *services.GroupService
	ParameterService *services.ParameterService
}

func NewHandler(serviceRepo repositories.ServiceRepository, groupRepo repositories.GroupRepository, paramRepo repositories.ParameterRepository) *Handler {
	groupService := services.NewGroupService(groupRepo)
	parameterService := services.NewParameterService(paramRepo)
	return &Handler{
		ServiceRepo:      serviceRepo,
		GroupRepo:        groupRepo,
		ParameterRepo:    paramRepo,
		GroupService:     groupService,
		ParameterService: parameterService,
	}
}
