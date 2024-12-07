package models

import (
	"time"
)

type Parameter struct {
	Code  string `gorm:"primaryKey" json:"code"`
	Title string `json:"title"`
}

type Service struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string      `json:"title"`
	Parameters []Parameter `gorm:"many2many:service_parameters;" json:"parameters"`
	GroupID    *uint       `gorm:"default:null" json:"group_id"`
	Group      *Group      `gorm:"foreignKey:GroupID" json:"group"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	ApprovedAt *time.Time  `json:"approved_at"`
}

type Group struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}
