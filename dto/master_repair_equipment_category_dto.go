package dto

type EquipmentCategoryManageBodyRequest struct {
	EquipmentNameCategory string `json:"equipment_name_category"`
	EquipmentCodeCategory string `json:"equipment_code_category"`
}

type EquipmentCategoryManageQueryRow struct {
	ID                    uint   `gorm:"id"`
	EquipmentNameCategory string `gorm:"equipment_name_category"`
	EquipmentCodeCategory string `gorm:"equipment_code_category"`
}

type EquipmentCategoryManageRespone struct {
	ID                    uint   `json:"id"`
	EquipmentNameCategory string `json:"equipmentNameCategory"`
	EquipmentCodeCategory string `json:"equipmentCodeCategory"`
}
