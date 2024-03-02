package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
	"github.com/gofiber/fiber/v2"
)

type masterIconController struct {
	repository repository.IMasterIconRepository
}

func NewMasterIconController() *masterIconController {
	return &masterIconController{
		repository: repository.NewMasterIconRepository(),
	}
}

func (s *masterIconController) IconCreate(c *fiber.Ctx) error {
	return c.Render("pages/admin/icon_create", fiber.Map{}, "layouts/admin/main")
}

func (s *masterIconController) IconAll(c *fiber.Ctx) error {
	icons := s.repository.GetIcons()

	var iconResp []dto.MasterIconRespone
	for _, icon := range icons {
		iconResp = append(iconResp, dto.MasterIconRespone(icon))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": iconResp})
}

func (s *masterIconController) IconSave(c *fiber.Ctx) error {
	var body dto.MasterIconBodyRequest
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := s.repository.CreateIcon(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"body": body})
}
