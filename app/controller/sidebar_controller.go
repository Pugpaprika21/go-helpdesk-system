package controller

import (
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type sidebarController struct {
	repository repository.ISidebarRepository
}

func NewSidebarController() *sidebarController {
	return &sidebarController{
		repository: repository.NewSidebarRepository(),
	}
}

func (s *sidebarController) SidebarMainCreate(c *fiber.Ctx) error {
	return c.Render("pages/admin/sidebar_main_create", fiber.Map{}, "layouts/admin/main")
}

func (s *sidebarController) SidebarMainSave(c *fiber.Ctx) error {
	var body dto.SideBarFormDataBodyRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	sidebarFile, _ := c.FormFile("sidebar_logo")

	fileExt := strings.Split(sidebarFile.Filename, ".")[1]
	sidebarLogo := fmt.Sprintf("%s.%s", strings.Replace(uuid.New().String(), "-", "", -1), fileExt)

	_, err := s.repository.CreateSidebarParent(body, sidebarLogo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create SidebarMain"})
	}

	c.SaveFile(sidebarFile, fmt.Sprintf("./public/uploads/%s", sidebarLogo))

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Sidebar saved successfully"})
}

func (s *sidebarController) SidebarCreate(c *fiber.Ctx) error {
	return c.Render("pages/admin/sidebar_menu_create", fiber.Map{}, "layouts/admin/main")
}

func (s *sidebarController) SidebarAll(c *fiber.Ctx) error {
	sidebars, err := s.repository.GetSidebars()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var sideBarResponses []dto.SideBarRespone
	for _, sidebar := range sidebars {
		var sidebarSubResponses []dto.SidebarSubRespone

		for _, sidebarSub := range sidebar.SidebarSub {
			sidebarSubResponses = append(sidebarSubResponses, dto.SidebarSubRespone(sidebarSub))
		}

		sidebarIcon := s.repository.GetSidebarIconName(sidebar.IconID)

		htmlTags := html.UnescapeString(sidebarIcon)

		sideBarResponses = append(sideBarResponses, dto.SideBarRespone{
			ID:                sidebar.ID,
			SidebarName:       sidebar.SidebarName,
			SidebarIcon:       htmlTags,
			SidebarTextColor:  sidebar.SidebarTextColor,
			SidebarBgColor:    sidebar.SidebarBgColor,
			SidebarSubRespone: sidebarSubResponses,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": sideBarResponses})
}

func (s *sidebarController) SidebarMenuSave(c *fiber.Ctx) error {
	var body dto.SideBarJsonBodyRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	id, err := s.repository.CreateSidebarChildMain(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create SidebarMain"})
	}

	s.repository.CreateSidebarChildSub(id, body)

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Sidebar saved successfully"})
}
