package models

import (
	"time"
)

type Parameter struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Service struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string      `json:"title"`
	Parameters []Parameter `gorm:"many2many:service_parameters;" json:"parameters"`
	ClassID    *uint       `gorm:"default:null" json:"class_id"`
	Class      *Class      `gorm:"foreignKey:ClassID" json:"class"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	ApprovedAt *time.Time  `json:"approved_at"`
}

type Class struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
