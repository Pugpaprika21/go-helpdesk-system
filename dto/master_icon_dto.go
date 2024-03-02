package dto

type MasterIconBodyRequest struct {
	IconName    string `json:"icon_name"`
	IconTagHtml string `json:"icon_tag_html"`
}

type MasterIconQueryRow struct {
	ID          uint   `gorm:"id"`
	IconName    string `gorm:"icon_name"`
	IconTagHtml string `gorm:"icon_tag_html"`
}

type MasterIconRespone struct {
	ID          uint   `json:"id"`
	IconName    string `json:"icon_name"`
	IconTagHtml string `json:"icon_tag_html"`
}
