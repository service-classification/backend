package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type GroupService struct {
	GroupRepository repositories.GroupRepository
}

func NewGroupService(groupRepo repositories.GroupRepository) *GroupService {
	return &GroupService{
		GroupRepository: groupRepo,
	}
}

type NewGroup struct {
	ID                uint     `json:"id" example:"3042"`
	Title             string   `json:"title"`
	AllowedParameters []string `json:"allowed_parameters" example:"mob_inet,fix_ctv,voice_fix"`
}

func (s *GroupService) CreateGroup(group NewGroup) (*models.Group, error) {
	model := &models.Group{
		ID:    group.ID,
		Title: group.Title,
	}
	err := s.GroupRepository.Create(model)
	if err != nil {
		return model, err
	}

	//todo add new group to the knowledge base

	return model, nil
}
