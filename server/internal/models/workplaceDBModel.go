package models

import "github.com/jackc/pgtype"

type WorkplaceModelDB struct {
	ID   int          `gorm:"type:int;autoIncrement"`
	Data pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null"`
}
