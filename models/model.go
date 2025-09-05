package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subs struct {
	gorm.Model
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	ServiceName string     `gorm:"not null" json:"service_name"`
	Price       int        `gorm:"not null" json:"price"`
	StartDate   MonthYear  `gorm:"type:date;not null" json:"start_date"`
	EndDate     *MonthYear `gorm:"type:date" json:"end_date,omitempty"`
}
