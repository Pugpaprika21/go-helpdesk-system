package dto

type RepairSaveBodyRequest struct {
	ID                  uint   `form:"id,omitempty"`
	UserID              uint   `form:"user_id,omitempty"`
	EquipmentRegisNum   string `form:"equipment_regis_num"`
	EquipmentCategoryID uint   `form:"equipment_category_id"`
	EquipmentCategory   string `form:"equipment_category"`
	EquipmentType       string `form:"equipment_type"`
	EquipmentBrand      string `form:"equipment_brand"`
	EquipmentDetail     string `form:"equipment_detail"`
}

type RepairQueryRow struct {
	ID                  uint   `gorm:"id"`
	UserID              uint   `gorm:"user_id"`
	EquipmentRegisNum   string `gorm:"equipment_regis_num"`
	EquipmentCategoryID uint   `gorm:"equipment_category_id"`
	EquipmentCategory   string `gorm:"equipment_category"`
	EquipmentType       string `gorm:"equipment_type"`
	EquipmentBrand      string `gorm:"equipment_brand"`
	EquipmentDetail     string `gorm:"equipment_detail"`
	EquipmentStatus     string `gorm:"equipment_status"`
}

type RepairEquipmentImgRespone struct {
	FileID                uint   `json:"fileId"`
	FileName              string `json:"fileName"`
	EquipmentFullPathImgs string `json:"equipmentFullPathImgs,omitempty"`
}

type RepairRespone struct {
	ID                    uint                        `json:"id"`
	UserID                uint                        `json:"userId,omitempty"`
	EquipmentRegisNum     string                      `json:"equipmentRegisNum"`
	EquipmentCategoryID   uint                        `json:"equipmentCategoryId"`
	EquipmentCategoryName string                      `json:"equipmentCategoryName"`
	EquipmentCategory     string                      `json:"equipmentCategory"`
	EquipmentType         string                      `json:"equipmentType"`
	EquipmentBrand        string                      `json:"equipmentBrand"`
	EquipmentStatus       string                      `json:"equipmentStatus"`
	EquipmentDetail       string                      `json:"equipmentDetail"`
	EquipmentFullPathImgs []RepairEquipmentImgRespone `json:"equipmentFullPathImgs,omitempty"`
	EquipmentSaveImgs     []FileStorageSystemRespone  `json:"equipmentSaveImgs"`
}
