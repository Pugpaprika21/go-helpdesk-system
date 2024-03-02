package models

import (
	"time"

	"gorm.io/gorm"
)

type MasterEquipmentCategoryManage struct {
	gorm.Model
	EquipmentNameCategory string    `gorm:"type:varchar(100); not null"`
	EquipmentCodeCategory string    `gorm:"type:varchar(200); not null"`
	CreatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
