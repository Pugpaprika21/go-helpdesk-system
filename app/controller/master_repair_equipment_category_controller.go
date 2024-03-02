package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
	"github.com/gofiber/fiber/v2"
)

type masterRepairEquipmentCategoryController struct {
	repository repository.IMasterRepairEquipmentCategory
}

func NewMasterRepairEquipmentCategoryController() *masterRepairEquipmentCategoryController {
	return &masterRepairEquipmentCategoryController{
		repository: repository.NewMasterRepairEquipmentCategory(),
	}
}

func (m *masterRepairEquipmentCategoryController) RepairEquipmentCategoryCreate(c *fiber.Ctx) error {
	return c.Render("pages/admin/master_equipment_category_manage", fiber.Map{}, "layouts/admin/main")
}

func (m *masterRepairEquipmentCategoryController) EquipmentCategorySave(c *fiber.Ctx) error {
	var body dto.EquipmentCategoryManageBodyRequest
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := m.repository.RepairEquipmentCategoryCreate(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": nil})
}

func (m *masterRepairEquipmentCategoryController) GetEquipmentCategoryList(c *fiber.Ctx) error {
	results := m.repository.GetRepairEquipmentCategoryAll()

	var equipments []dto.EquipmentCategoryManageRespone
	for _, result := range results {
		equipments = append(equipments, dto.EquipmentCategoryManageRespone{
			ID:                    result.ID,
			EquipmentNameCategory: result.EquipmentNameCategory,
			EquipmentCodeCategory: result.EquipmentCodeCategory,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": equipments})
}
