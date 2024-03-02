package repository

import (
	"time"

	"github.com/Pugpaprika21/go-fiber/app/db"
	"github.com/Pugpaprika21/go-fiber/app/models"
	"github.com/Pugpaprika21/go-fiber/dto"
	"gorm.io/gorm"
)

type IRepairRepository interface {
	CreateRepair(body dto.RepairSaveBodyRequest) (uint, error)
	CreateImageRepair(repairID uint, equipmentImg string, fileSize int64, fileExt string) (uint, error)
	GetRepairList(id uint) ([]*dto.RepairQueryRow, error)
	GetRepairListByUserID(userID uint) ([]*dto.RepairQueryRow, error)
	GetRepairImgByRepairID(repairID uint) ([]*dto.FileStorageSystemQueryRow, error)
	GetEquipmentCategoryName(equipmentCategoryID uint) string
	RepairUpdateByID(id uint, body dto.RepairSaveBodyRequest) (uint, error)
	RepairUpdateImgByRepairID(repairID uint, equipmentImg string, fileSize int64, fileExt string) (uint, error)
	DeleteRepairByID(id uint) error
	DeleteRepairImgByID(imgID uint) error
}

type repairRepository struct {
	DB *gorm.DB
}

func NewRepairRepository() *repairRepository {
	return &repairRepository{
		DB: db.Conn,
	}
}

func (r *repairRepository) CreateRepair(body dto.RepairSaveBodyRequest) (uint, error) {
	repair := &models.Repair{
		EquipmentRegisNum:   body.EquipmentRegisNum,
		EquipmentCategoryID: body.EquipmentCategoryID,
		EquipmentCategory:   body.EquipmentCategory,
		EquipmentType:       body.EquipmentType,
		EquipmentBrand:      body.EquipmentBrand,
		EquipmentDetail:     body.EquipmentDetail,
		EquipmentStatus:     "progress",
		UserID:              1,
		CreatedAt:           time.Now(),
	}
	result := r.DB.Model(&models.Repair{}).Create(repair)
	if result.Error != nil {
		return 0, result.Error
	}

	return repair.ID, nil
}

func (r *repairRepository) CreateImageRepair(repairID uint, equipmentImg string, fileSize int64, fileExt string) (uint, error) {
	imgRepair := &models.FileStorageSystem{
		FileName:      equipmentImg,
		FileSize:      fileSize,
		FileType:      fileExt,
		FileExtension: "",
		Content:       "",
		RefID:         repairID,
		RefTable:      "repairs",
		RefField:      "id",
		CreatedAt:     time.Now(),
	}

	result := r.DB.Model(&models.FileStorageSystem{}).Create(imgRepair)
	if result.Error != nil {
		return 0, result.Error
	}

	return imgRepair.ID, nil
}

func (r *repairRepository) GetRepairList(id uint) ([]*dto.RepairQueryRow, error) {
	var repairs []*dto.RepairQueryRow
	query := r.DB.Model(&models.Repair{}).Order("created_at DESC")
	if id != 0 {
		query = query.Where("id = ?", id)
	}

	if err := query.Find(&repairs).Error; err != nil {
		return nil, err
	}
	return repairs, nil
}

func (r *repairRepository) GetRepairListByUserID(userID uint) ([]*dto.RepairQueryRow, error) {
	var repairs []*dto.RepairQueryRow
	if err := r.DB.Model(&models.Repair{}).Where("user_id = ?", userID).Find(&repairs).Error; err != nil {
		return nil, err
	}
	return repairs, nil
}

func (r *repairRepository) GetRepairImgByRepairID(repairID uint) ([]*dto.FileStorageSystemQueryRow, error) {
	var repairImgs []*dto.FileStorageSystemQueryRow
	if err := r.DB.Model(&models.FileStorageSystem{}).Where("ref_id = ? AND ref_table = ?", repairID, "repairs").Find(&repairImgs).Error; err != nil {
		return nil, err
	}
	return repairImgs, nil
}

func (r *repairRepository) GetEquipmentCategoryName(equipmentCategoryID uint) string {
	var result dto.EquipmentCategoryManageQueryRow
	r.DB.Model(&models.MasterEquipmentCategoryManage{}).Where("id = ?", equipmentCategoryID).First(&result)

	return result.EquipmentNameCategory
}

func (r *repairRepository) RepairUpdateByID(id uint, body dto.RepairSaveBodyRequest) (uint, error) {
	repair := &models.Repair{
		EquipmentRegisNum:   body.EquipmentRegisNum,
		EquipmentCategoryID: body.EquipmentCategoryID,
		EquipmentCategory:   body.EquipmentCategory,
		EquipmentType:       body.EquipmentType,
		EquipmentBrand:      body.EquipmentBrand,
		EquipmentDetail:     body.EquipmentDetail,
		EquipmentStatus:     "progress",
	}

	result := r.DB.Model(&models.Repair{}).Where("id = ?", id).Updates(repair)
	if result.Error != nil {
		return 0, result.Error
	}

	return repair.ID, nil
}

func (r *repairRepository) RepairUpdateImgByRepairID(repairID uint, equipmentImg string, fileSize int64, fileExt string) (uint, error) {
	return 0, nil
}

func (r *repairRepository) DeleteRepairByID(id uint) error {
	return r.DB.Where("id = ?", id).Unscoped().Delete(&models.Repair{}).Error
}

func (r *repairRepository) DeleteRepairImgByID(imgID uint) error {
	return r.DB.Where("id = ?", imgID).Unscoped().Delete(&models.FileStorageSystem{}).Error
}
