package models

import (
	"time"
)

type Parameter struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	New       bool      `json:"new"`
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
	New       bool      `json:"new"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ClassView struct {
	ID                uint     `json:"id" example:"3042"`
	Title             string   `json:"title"`
	AllowedParameters []string `json:"allowed_parameters" example:"mob_inet,fix_ctv,voice_fix"`
}

type ParameterView struct {
	ID                      string   `json:"id" example:"fix_ctv" required:"true,alphanum"`
	Title                   string   `json:"title"`
	AllowedClasses          []uint   `json:"allowed_classes" example:"1,1033,3023"`
	ContradictionParameters []string `json:"contradiction_parameters" example:"mob_inet,fix_ctv,voice_fix"`
}
