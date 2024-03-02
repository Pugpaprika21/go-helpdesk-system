package models

import (
	"time"

	"gorm.io/gorm"
)

type MasterIcon struct {
	gorm.Model
	IconName    string    `gorm:"type:varchar(100); not null"`
	IconTagHtml string    `gorm:"type:varchar(200); not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
