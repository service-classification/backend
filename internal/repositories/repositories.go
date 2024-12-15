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
	FindByParameterID(parameterID string) ([]models.Service, error)
	FindByClassID(id uint) ([]models.Service, error)
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
	return r.db.Model(service).Omit("CreatedAt").Updates(service).Error
}

func (r *serviceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Service{}, id).Error
}

func (r *serviceRepository) GetByID(id uint) (*models.Service, error) {
	var service models.Service
	err := r.db.Preload("Parameters").Preload("Class").First(&service, id).Error
	return &service, err
}

func (r *serviceRepository) List(offset, limit int) ([]models.Service, error) {
	var services []models.Service
	err := r.db.
		Preload("Class").
		Offset(offset).
		Limit(limit).
		Order("approved_at IS NOT NULL, approved_at DESC").
		Find(&services).
		Error
	return services, err
}

func (r *serviceRepository) FindByParameterID(parameterID string) ([]models.Service, error) {
	var services []models.Service
	err := r.db.
		Preload("Parameters").
		Joins("JOIN service_parameters ON services.id = service_parameters.service_id").
		Joins("JOIN parameters ON service_parameters.parameter_id = parameters.id").
		Where("parameters.id = ?", parameterID).
		Find(&services).Error
	return services, err
}

func (r *serviceRepository) FindByClassID(id uint) ([]models.Service, error) {
	var services []models.Service
	err := r.db.
		Preload("Class").
		Where("class_id = ?", id).
		Find(&services).Error
	return services, err
}

type ClassRepository interface {
	GetByID(id uint) (*models.Class, error)
	List(offset, limit int) ([]models.Class, error)
	Update(class *models.Class) error
	Create(class *models.Class) error
	Delete(u uint) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db}
}

func (r *classRepository) GetByID(id uint) (*models.Class, error) {
	var class models.Class
	err := r.db.First(&class, id).Error
	return &class, err
}

func (r *classRepository) List(offset, limit int) ([]models.Class, error) {
	var classes []models.Class
	err := r.db.Offset(offset).Limit(limit).Find(&classes).Error
	return classes, err
}

func (r *classRepository) Create(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *classRepository) Update(class *models.Class) error {
	return r.db.Model(class).Omit("CreatedAt").Updates(class).Error
}

func (r *classRepository) Delete(u uint) error {
	return r.db.Delete(&models.Class{}, u).Error
}

type ParameterRepository interface {
	Create(parameter *models.Parameter) error
	Update(parameter *models.Parameter) error
	Delete(code string) error
	GetByID(code string) (*models.Parameter, error)
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
	return r.db.Model(parameter).Omit("CreatedAt").Updates(parameter).Error
}

func (r *parameterRepository) Delete(code string) error {
	return r.db.Delete(&models.Parameter{}, "id = ?", code).Error
}

func (r *parameterRepository) GetByID(code string) (*models.Parameter, error) {
	var parameter models.Parameter
	err := r.db.First(&parameter, "id = ?", code).Error
	return &parameter, err
}

func (r *parameterRepository) List(offset, limit int) ([]models.Parameter, error) {
	var parameters []models.Parameter
	err := r.db.Offset(offset).Limit(limit).Find(&parameters).Error
	return parameters, err
}
