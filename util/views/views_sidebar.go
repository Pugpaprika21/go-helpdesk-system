package util

import (
	"html"

	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
)

type sidebar struct {
	repository repository.ISidebarRepository
}

func NewSidebar() *sidebar {
	return &sidebar{
		repository: repository.NewSidebarRepository(),
	}
}

func (s *sidebar) renderAll() []dto.SideBarRespone {
	sidebars, _ := s.repository.GetSidebars()

	var sideBarResponses []dto.SideBarRespone
	for _, sidebar := range sidebars {
		var sidebarSubResponses []dto.SidebarSubRespone

		for _, sidebarSub := range sidebar.SidebarSub {
			sidebarSubResponses = append(sidebarSubResponses, dto.SidebarSubRespone(sidebarSub))
		}

		sidebarIcon := s.repository.GetSidebarIconName(sidebar.IconID)

		decodedIcon := html.UnescapeString(sidebarIcon)

		sideBarResponses = append(sideBarResponses, dto.SideBarRespone{
			ID:                sidebar.ID,
			SidebarName:       sidebar.SidebarName,
			SidebarIcon:       string(decodedIcon),
			SidebarTextColor:  sidebar.SidebarTextColor,
			SidebarBgColor:    sidebar.SidebarBgColor,
			SidebarSubRespone: sidebarSubResponses,
		})
	}

	return sideBarResponses
}
