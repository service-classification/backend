package repositories

import (
	"backend/internal/models"
	"errors"

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
	err := r.db.First(&service, id).Error
	return &service, err
}

func (r *serviceRepository) List(offset, limit int) ([]models.Service, error) {
	var services []models.Service
	err := r.db.Offset(offset).Limit(limit).Find(&services).Error
	return services, err
}

type ClassifiedServiceRepository interface {
	AssignGroupToService(serviceID uint, groupID uint) error
	RemoveGroupFromService(serviceID uint, groupID uint) error
	GetGroupsByServiceID(serviceID uint) ([]models.ClassifiedService, error)
}

type classifiedServiceRepository struct {
	db *gorm.DB
}

func NewClassifiedServiceRepository(db *gorm.DB) ClassifiedServiceRepository {
	return &classifiedServiceRepository{db}
}

func (r *classifiedServiceRepository) AssignGroupToService(serviceID uint, groupID uint) error {
	// Check if the association already exists
	var cs models.ClassifiedService
	err := r.db.Where("service_id = ? AND group_id = ?", serviceID, groupID).First(&cs).Error
	if err == nil {
		// Association already exists
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// An unexpected error occurred
		return err
	}

	// Create the association
	cs = models.ClassifiedService{
		ServiceID: serviceID,
		GroupID:   groupID,
	}
	return r.db.Create(&cs).Error
}

func (r *classifiedServiceRepository) RemoveGroupFromService(serviceID uint, groupID uint) error {
	return r.db.Where("service_id = ? AND group_id = ?", serviceID, groupID).Delete(&models.ClassifiedService{}).Error
}

func (r *classifiedServiceRepository) GetGroupsByServiceID(serviceID uint) ([]models.ClassifiedService, error) {
	var csList []models.ClassifiedService
	err := r.db.Where("service_id = ?", serviceID).Find(&csList).Error
	return csList, err
}

type GroupRepository interface {
	GetByID(id uint) (*models.Group, error)
	List(offset, limit int) ([]models.Group, error)
	// Additional methods as needed
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
