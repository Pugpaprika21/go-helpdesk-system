package repository

import (
	"time"

	"github.com/Pugpaprika21/go-fiber/app/db"
	"github.com/Pugpaprika21/go-fiber/app/models"
	"github.com/Pugpaprika21/go-fiber/dto"
	"gorm.io/gorm"
)

type IMasterRepairEquipmentCategory interface {
	RepairEquipmentCategoryCreate(body dto.EquipmentCategoryManageBodyRequest) (uint, error)
	GetRepairEquipmentCategoryAll() []*dto.EquipmentCategoryManageQueryRow
}

type masterRepairEquipmentCategory struct {
	DB *gorm.DB
}

func NewMasterRepairEquipmentCategory() *masterRepairEquipmentCategory {
	return &masterRepairEquipmentCategory{
		DB: db.Conn,
	}
}

func (m *masterRepairEquipmentCategory) RepairEquipmentCategoryCreate(body dto.EquipmentCategoryManageBodyRequest) (uint, error) {
	repair := &models.MasterEquipmentCategoryManage{
		EquipmentNameCategory: body.EquipmentNameCategory,
		EquipmentCodeCategory: body.EquipmentCodeCategory,
		CreatedAt:             time.Now(),
	}

	result := m.DB.Model(&models.MasterEquipmentCategoryManage{}).Create(repair)
	if result.Error != nil {
		return 0, result.Error
	}

	return repair.ID, nil
}

func (m *masterRepairEquipmentCategory) GetRepairEquipmentCategoryAll() []*dto.EquipmentCategoryManageQueryRow {
	var equipments []*dto.EquipmentCategoryManageQueryRow
	m.DB.Model(&models.MasterEquipmentCategoryManage{}).Find(&equipments)

	return equipments
}
