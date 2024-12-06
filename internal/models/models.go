package models

type Parameter struct {
	Code  string `gorm:"primaryKey" json:"code"`
	Title string `json:"title"`
}

type Service struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Parameters  []Parameter `gorm:"many2many:service_parameters;" json:"parameters"`
}

type Group struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}

type ClassifiedService struct {
	ServiceID uint    `gorm:"primaryKey" json:"service_id"`
	GroupID   uint    `gorm:"primaryKey" json:"group_id"`
	Service   Service `gorm:"foreignKey:ServiceID" json:"-"`
	Group     Group   `gorm:"foreignKey:GroupID" json:"-"`
}
