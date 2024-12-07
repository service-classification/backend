package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(service *models.Service) error
	Update(service *models.Service) error
	Delete(id uint) error
	GetByID(id uint) (*models.Service, error)
	List(offset, limit int) ([]models.Service, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) Create(service *models.Service) error {
	return r.db.Create(service).Error
}

func (r *serviceRepository) Update(service *models.Service) error {
	return r.db.Save(service).Error
}

func (r *serviceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Service{}, id).Error
}

func (r *serviceRepository) GetByID(id uint) (*models.Service, error) {
	var service models.Service
	err := r.db.Preload("Parameters").Preload("Group").First(&service, id).Error
	return &service, err
}

func (r *serviceRepository) List(offset, limit int) ([]models.Service, error) {
	var services []models.Service
	err := r.db.Preload("Parameters").Preload("Group").Offset(offset).Limit(limit).Find(&services).Error
	return services, err
}

type GroupRepository interface {
	GetByID(id uint) (*models.Group, error)
	List(offset, limit int) ([]models.Group, error)
	Create(group *models.Group) error
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db}
}

func (r *groupRepository) GetByID(id uint) (*models.Group, error) {
	var group models.Group
	err := r.db.First(&group, id).Error
	return &group, err
}

func (r *groupRepository) List(offset, limit int) ([]models.Group, error) {
	var groups []models.Group
	err := r.db.Offset(offset).Limit(limit).Find(&groups).Error
	return groups, err
}

func (r *groupRepository) Create(group *models.Group) error {
	return r.db.Create(group).Error
}

type ParameterRepository interface {
	Create(parameter *models.Parameter) error
	Update(parameter *models.Parameter) error
	Delete(code string) error
	GetByCode(code string) (*models.Parameter, error)
	List(offset, limit int) ([]models.Parameter, error)
}

type parameterRepository struct {
	db *gorm.DB
}

func NewParameterRepository(db *gorm.DB) ParameterRepository {
	return &parameterRepository{db}
}

func (r *parameterRepository) Create(parameter *models.Parameter) error {
	return r.db.Create(parameter).Error
}

func (r *parameterRepository) Update(parameter *models.Parameter) error {
	return r.db.Save(parameter).Error
}

func (r *parameterRepository) Delete(code string) error {
	return r.db.Delete(&models.Parameter{}, "code = ?", code).Error
}

func (r *parameterRepository) GetByCode(code string) (*models.Parameter, error) {
	var parameter models.Parameter
	err := r.db.First(&parameter, "code = ?", code).Error
	return &parameter, err
}

func (r *parameterRepository) List(offset, limit int) ([]models.Parameter, error) {
	var parameters []models.Parameter
	err := r.db.Offset(offset).Limit(limit).Find(&parameters).Error
	return parameters, err
}
