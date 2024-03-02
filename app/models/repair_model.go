package models

import (
	"time"

	"gorm.io/gorm"
)

type Repair struct {
	gorm.Model
	EquipmentRegisNum             string `gorm:"type:varchar(200); not null"`
	EquipmentCategoryID           uint
	MasterEquipmentCategoryManage MasterEquipmentCategoryManage `gorm:"foreignKey:EquipmentCategoryID"`
	EquipmentCategory             string                        `gorm:"type:varchar(200); not null"`
	EquipmentType                 string                        `gorm:"type:varchar(200); not null"`
	EquipmentBrand                string                        `gorm:"type:varchar(200); not null"`
	EquipmentDetail               string                        `gorm:"type:varchar(200); not null"`
	EquipmentStatus               string                        `gorm:"type:ENUM('progress', 'succeed'); not null"`
	UserID                        uint
	User                          User      `gorm:"foreignKey:UserID"`
	CreatedAt                     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
