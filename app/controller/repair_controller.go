package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type repairController struct {
	repository repository.IRepairRepository
}

func NewRepairController() *repairController {
	return &repairController{
		repository: repository.NewRepairRepository(),
	}
}

func (r *repairController) RepairFormSave(c *fiber.Ctx) error {
	return c.Render("pages/user/repair_form_save", fiber.Map{}, "layouts/admin/main")
}

func (r *repairController) RepairSave(c *fiber.Ctx) error {
	var body dto.RepairSaveBodyRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	repairID, err := r.repository.CreateRepair(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create"})
	}

	fileForm, _ := c.MultipartForm()
	for _, file := range fileForm.File["equipment_save_imgs"] {
		fileExt := strings.Split(file.Filename, ".")[1]
		equipmentImg := fmt.Sprintf("%s.%s", strings.Replace(uuid.New().String(), "-", "", -1), fileExt)

		_, err := r.repository.CreateImageRepair(repairID, equipmentImg, file.Size, fileExt)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", equipmentImg))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": nil})
}

func (r *repairController) RepairList(c *fiber.Ctx) error {
	return c.Render("pages/user/repair_list", fiber.Map{}, "layouts/admin/main")
}

func (r *repairController) RepairGetList(c *fiber.Ctx) error {
	results, _ := r.repository.GetRepairList(0)

	var resp []dto.RepairRespone
	for _, result := range results {
		equipmentCategoryName := r.repository.GetEquipmentCategoryName(result.EquipmentCategoryID)
		resp = append(resp, dto.RepairRespone{
			ID:                    result.ID,
			UserID:                result.UserID,
			EquipmentRegisNum:     result.EquipmentRegisNum,
			EquipmentCategoryID:   result.EquipmentCategoryID,
			EquipmentCategoryName: equipmentCategoryName,
			EquipmentCategory:     result.EquipmentCategory,
			EquipmentType:         result.EquipmentType,
			EquipmentBrand:        result.EquipmentBrand,
			EquipmentDetail:       result.EquipmentDetail,
			EquipmentStatus:       result.EquipmentStatus,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": resp})
}

func (r *repairController) RepairDelete(c *fiber.Ctx) error {
	repairID, _ := strconv.ParseUint(c.Params("repair_id"), 10, 64)

	err := r.repository.DeleteRepairByID(uint(repairID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	repairImgs, _ := r.repository.GetRepairImgByRepairID(uint(repairID))
	for _, img := range repairImgs {
		r.repository.DeleteRepairImgByID(img.ID)
		filePath := fmt.Sprintf("./public/uploads/%s", img.FileName)
		os.Remove(filePath)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": ""})
}

func (r *repairController) RepairFormEdit(c *fiber.Ctx) error {
	return c.Render("pages/user/repair_form_edit", fiber.Map{}, "layouts/admin/main")
}

func (r *repairController) RepairEdit(c *fiber.Ctx) error {
	repairID, _ := strconv.ParseUint(c.Params("repair_id"), 10, 64)

	results, err := r.repository.GetRepairList(uint(repairID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var repairResp []dto.RepairRespone
	for _, result := range results {
		var equipmentFullPathImgs []dto.RepairEquipmentImgRespone
		repairImgs, _ := r.repository.GetRepairImgByRepairID(result.ID)
		for _, repairImg := range repairImgs {
			equipmentFullPathImgs = append(equipmentFullPathImgs, dto.RepairEquipmentImgRespone{
				FileID:                repairImg.ID,
				FileName:              repairImg.FileName,
				EquipmentFullPathImgs: os.Getenv("APP_URL") + "api/repair/repair_attached_files?name=" + repairImg.FileName,
			})
		}

		repairResp = append(repairResp, dto.RepairRespone{
			ID:                    result.ID,
			EquipmentRegisNum:     result.EquipmentRegisNum,
			EquipmentCategoryID:   result.EquipmentCategoryID,
			EquipmentCategory:     result.EquipmentCategory,
			EquipmentType:         result.EquipmentType,
			EquipmentBrand:        result.EquipmentBrand,
			EquipmentDetail:       result.EquipmentDetail,
			EquipmentStatus:       result.EquipmentStatus,
			EquipmentFullPathImgs: equipmentFullPathImgs,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": repairResp[0]})
}

func (r *repairController) RepairAttachedFiles(c *fiber.Ctx) error {
	fileName := c.Query("name")
	filePath := fmt.Sprintf("./public/uploads/%s", fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}
	return c.SendFile(filePath)
}

func (r *repairController) RepairDeleteAttachedFile(c *fiber.Ctx) error {
	fileID, _ := strconv.ParseUint(c.Params("file_id"), 10, 64)
	fileName := c.Params("file_name")

	err := r.repository.DeleteRepairImgByID(uint(fileID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	filePath := fmt.Sprintf("./public/uploads/%s", fileName)
	os.Remove(filePath)

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": fileID, "fileName": fileName})
}

func (r *repairController) RepairEditSave(c *fiber.Ctx) error {
	repairID, _ := strconv.ParseUint(c.Params("repair_id"), 10, 64)

	var body dto.RepairSaveBodyRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := r.repository.RepairUpdateByID(uint(repairID), body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create"})
	}

	fileForm, _ := c.MultipartForm()
	for _, file := range fileForm.File["equipment_save_imgs"] {
		fileExt := strings.Split(file.Filename, ".")[1]
		equipmentImg := fmt.Sprintf("%s.%s", strings.Replace(uuid.New().String(), "-", "", -1), fileExt)

		r.repository.CreateImageRepair(uint(repairID), equipmentImg, file.Size, fileExt)

		c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", equipmentImg))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": nil})
}
