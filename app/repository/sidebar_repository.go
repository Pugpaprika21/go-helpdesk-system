package repository

import (
	"github.com/Pugpaprika21/go-fiber/app/db"
	"github.com/Pugpaprika21/go-fiber/app/models"
	"github.com/Pugpaprika21/go-fiber/dto"
	"gorm.io/gorm"
)

type ISidebarRepository interface {
	CreateSidebarParent(body dto.SideBarFormDataBodyRequest, sidebarLogo string) (uint, error)
	CreateSidebarChildMain(body dto.SideBarJsonBodyRequest) (uint, error)
	CreateSidebarChildSub(sidebarMainID uint, body dto.SideBarJsonBodyRequest) (uint, error)
	GetSidebars() ([]dto.SidebarChildMainQueryRow, error)
	GetSidebarIconName(id uint) string
	GetSidebarLogoLastCreate()
}

type sidebarRepository struct {
	DB *gorm.DB
}

func NewSidebarRepository() *sidebarRepository {
	return &sidebarRepository{
		DB: db.Conn,
	}
}

func (s *sidebarRepository) CreateSidebarParent(body dto.SideBarFormDataBodyRequest, sidebarLogo string) (uint, error) {
	sidebarParent := &models.SidebarParent{
		SidebarLogo:      sidebarLogo,
		SidebarBgColor:   body.SidebarBgColor,
		SidebarTextColor: body.SidebarTextColor,
	}

	result := s.DB.Model(&models.SidebarParent{}).Create(sidebarParent)
	if result.Error != nil {
		return 0, result.Error
	}

	return sidebarParent.ID, nil
}

func (s *sidebarRepository) CreateSidebarChildMain(body dto.SideBarJsonBodyRequest) (uint, error) {
	sidebarMain := &models.SidebarChildMain{
		SidebarName: body.SidebarName,
		IconID:      body.IconID,
	}

	result := s.DB.Model(&models.SidebarChildMain{}).Create(sidebarMain)
	if result.Error != nil {
		return 0, result.Error
	}

	return sidebarMain.ID, nil
}

func (s *sidebarRepository) CreateSidebarChildSub(sidebarMainID uint, body dto.SideBarJsonBodyRequest) (uint, error) {
	for i := range body.SidebarSubNames {
		sidebarSub := &models.SidebarChildSub{
			SidebarMainID:   sidebarMainID,
			SidebarSubNames: body.SidebarSubNames[i],
			SidebarSubUrls:  body.SidebarSubUrls[i],
		}

		result := s.DB.Model(&models.SidebarChildSub{}).Create(sidebarSub)
		if result.Error != nil {
			return 0, result.Error
		}
	}

	return sidebarMainID, nil
}

func (s *sidebarRepository) GetSidebars() ([]dto.SidebarChildMainQueryRow, error) {
	var sidebarMains []models.SidebarChildMain
	if err := s.DB.Preload("SidebarChildSub").Order("order_num DESC").Find(&sidebarMains).Error; err != nil {
		return nil, err
	}

	var SidebarChildMainQueryRows []dto.SidebarChildMainQueryRow
	for _, sidebarMain := range sidebarMains {
		var sidebarSubQueryRows []dto.SidebarChildSubQueryRow
		for _, sidebarSub := range sidebarMain.SidebarChildSub {
			sub := dto.SidebarChildSubQueryRow{
				ID:              sidebarSub.ID,
				SidebarMainID:   sidebarSub.SidebarMainID,
				SidebarSubNames: sidebarSub.SidebarSubNames,
				SidebarSubUrls:  sidebarSub.SidebarSubUrls,
			}
			sidebarSubQueryRows = append(sidebarSubQueryRows, sub)
		}

		main := dto.SidebarChildMainQueryRow{
			ID:          sidebarMain.ID,
			SidebarName: sidebarMain.SidebarName,
			IconID:      sidebarMain.IconID,
			SidebarSub:  sidebarSubQueryRows,
		}
		SidebarChildMainQueryRows = append(SidebarChildMainQueryRows, main)
	}

	return SidebarChildMainQueryRows, nil
}

func (s *sidebarRepository) GetSidebarIconName(id uint) string {
	var iconResult dto.MasterIconQueryRow
	s.DB.Model(&models.MasterIcon{}).Where("id = ?", id).First(&iconResult)

	return iconResult.IconTagHtml
}

func (s *sidebarRepository) GetSidebarLogoLastCreate() {}
