package handlers

import (
	"backend/internal/repositories"
)

type Handler struct {
	ServiceRepo   repositories.ServiceRepository
	GroupRepo     repositories.GroupRepository
	ParameterRepo repositories.ParameterRepository
}

func NewHandler(serviceRepo repositories.ServiceRepository, groupRepo repositories.GroupRepository, paramRepo repositories.ParameterRepository) *Handler {
	return &Handler{
		ServiceRepo:   serviceRepo,
		GroupRepo:     groupRepo,
		ParameterRepo: paramRepo,
	}
}
