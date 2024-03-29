package migrtion

import (
	"github.com/Pugpaprika21/go-fiber/app/db"
	"github.com/Pugpaprika21/go-fiber/app/models"
)

func Run() {
	db.Conn.AutoMigrate(
		&models.User{},
		&models.SidebarParent{},
		&models.SidebarChildMain{},
		&models.SidebarChildSub{},
		&models.MasterIcon{},
		&models.MasterEquipmentCategoryManage{},
		&models.Repair{},
		&models.FileStorageSystem{},
	)
}
