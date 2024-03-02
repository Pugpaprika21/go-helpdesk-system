package models

import (
	"time"

	"gorm.io/gorm"
)

type SidebarParent struct {
	gorm.Model
	SidebarLogo      string    `gorm:"type:varchar(200); not null"`
	SidebarBgColor   string    `gorm:"type:varchar(100); not null"`
	SidebarTextColor string    `gorm:"type:varchar(50); not null"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type SidebarChildMain struct {
	gorm.Model
	IconID          uint
	SidebarName     string            `gorm:"type:varchar(100); not null"`
	MasterIcon      MasterIcon        `gorm:"foreignKey:IconID"`
	SidebarChildSub []SidebarChildSub `gorm:"foreignKey:SidebarMainID"`
	OrderNum        uint
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type SidebarChildSub struct {
	gorm.Model
	SidebarMainID   uint      `gorm:"not null"`
	SidebarSubNames string    `gorm:"type:varchar(100); not null"`
	SidebarSubUrls  string    `gorm:"type:varchar(100); not null"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
