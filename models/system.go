package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type System struct {
	gorm.Model

	Id      uuid.UUID `json:"ID" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Key     string    `json:"key"`
	Expired int64     `json:"expired" gorm:"default:EXTRACT(epoch FROM now())::bigint + 300"`
}

func (System) TableName() string {
	return "systems"
}
