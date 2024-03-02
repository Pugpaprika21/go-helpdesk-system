package util

import (
	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
)

type masterIcon struct {
	repository repository.IMasterIconRepository
}

func NewMasterIcon() *masterIcon {
	return &masterIcon{
		repository: repository.NewMasterIconRepository(),
	}
}

func (m *masterIcon) renderAll() []dto.MasterIconRespone {
	iconResult := m.repository.GetIcons()

	var icons []dto.MasterIconRespone
	for _, icon := range iconResult {
		icons = append(icons, dto.MasterIconRespone(icon))
	}
	return icons
}
