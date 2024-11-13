package models

import (
	"database/sql/driver"
	"errors"

	"gorm.io/gorm"
)

// Parameter enum
type Parameter string

const (
	MobInet             Parameter = "mob_inet"
	FixInet             Parameter = "fix_inet"
	FixCtv              Parameter = "fix_ctv"
	FixIctv             Parameter = "fix_ictv"
	VoiceMob            Parameter = "voice_mob"
	VoiceFix            Parameter = "voice_fix"
	Sms                 Parameter = "sms"
	Csd                 Parameter = "csd"
	Iot                 Parameter = "iot"
	Mms                 Parameter = "mms"
	Roaming             Parameter = "roaming"
	Mg                  Parameter = "mg"
	Mn                  Parameter = "mn"
	Mts                 Parameter = "mts"
	Conc                Parameter = "conc"
	FixOp               Parameter = "fix_op"
	VsrRoam             Parameter = "vsr_roam"
	NationalRoam        Parameter = "national_roam"
	MnRoam              Parameter = "mn_roam"
	VoiceAp             Parameter = "voice_ap"
	VoiceFee            Parameter = "voice_fee"
	PeriodService       Parameter = "period_service"
	OneTimeService      Parameter = "one_time_service"
	DopService          Parameter = "dop_service"
	Content             Parameter = "content"
	ServicesService     Parameter = "services_service"
	KeoSale             Parameter = "keo_sale"
	Discount            Parameter = "discount"
	OnlyInbound         Parameter = "only_inbound"
	SmsA2p              Parameter = "sms_a2p"
	SmsGross            Parameter = "sms_gross"
	Skoring             Parameter = "skoring"
	OtherService        Parameter = "other_service"
	VoiceMail           Parameter = "voice_mail"
	Geo                 Parameter = "geo"
	EpForNumber         Parameter = "ep_for_number"
	EpForLine           Parameter = "ep_for_line"
	OneTimeFeeForNumber Parameter = "one_time_fee_for_number"
	EquipmentRent       Parameter = "equipment rent"
	AddPackage          Parameter = "add_package"
)

var AllParameters = []Parameter{
	MobInet, FixInet, FixCtv, FixIctv, VoiceMob, VoiceFix, Sms, Csd, Iot, Mms,
	Roaming, Mg, Mn, Mts, Conc, FixOp, VsrRoam, NationalRoam, MnRoam, VoiceAp,
	VoiceFee, PeriodService, OneTimeService, DopService, Content, ServicesService,
	KeoSale, Discount, OnlyInbound, SmsA2p, SmsGross, Skoring, OtherService,
	VoiceMail, Geo, EpForNumber, EpForLine, OneTimeFeeForNumber, EquipmentRent,
	AddPackage,
}

// Implement the Scanner and Valuer interfaces for Parameter slice to store it in PostgreSQL

func (p *Parameter) Scan(value interface{}) error {
	*p = Parameter(value.(string))
	return nil
}

func (p Parameter) Value() (driver.Value, error) {
	return string(p), nil
}

type Service struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Parameters  []Parameter `gorm:"type:text[]" json:"parameters"`
}

func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	return s.validateParameters()
}

func (s *Service) BeforeUpdate(tx *gorm.DB) (err error) {
	return s.validateParameters()
}

func (s *Service) validateParameters() error {
	for _, param := range s.Parameters {
		if !isValidParameter(param) {
			return errors.New("invalid parameter: " + string(param))
		}
	}
	return nil
}

func isValidParameter(param Parameter) bool {
	for _, p := range AllParameters {
		if p == param {
			return true
		}
	}
	return false
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
