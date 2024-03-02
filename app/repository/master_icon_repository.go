package repository

import (
	"github.com/Pugpaprika21/go-fiber/app/db"
	"github.com/Pugpaprika21/go-fiber/app/models"
	"github.com/Pugpaprika21/go-fiber/dto"
	"gorm.io/gorm"
)

type IMasterIconRepository interface {
	CreateIcon(body dto.MasterIconBodyRequest) (uint, error)
	GetIcons() []dto.MasterIconQueryRow
}

type msterIconRepository struct {
	DB *gorm.DB
}

func NewMasterIconRepository() *msterIconRepository {
	return &msterIconRepository{
		DB: db.Conn,
	}
}

func (m *msterIconRepository) CreateIcon(body dto.MasterIconBodyRequest) (uint, error) {
	icon := &models.MasterIcon{
		IconName:    body.IconName,
		IconTagHtml: body.IconTagHtml,
	}

	result := m.DB.Model(&models.MasterIcon{}).Create(icon)
	if result.Error != nil {
		return 0, result.Error
	}

	return icon.ID, nil
}

func (m *msterIconRepository) GetIcons() []dto.MasterIconQueryRow {
	var icons []dto.MasterIconQueryRow
	m.DB.Model(&models.MasterIcon{}).Find(&icons)

	return icons
}
