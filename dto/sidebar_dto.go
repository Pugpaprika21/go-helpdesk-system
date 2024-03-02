package dto

type SideBarFormDataBodyRequest struct {
	SidebarName      string   `form:"sidebar_name"`
	IconID           uint     `form:"icon_id"`
	SidebarTextColor string   `form:"sidebar_text_color"`
	SidebarBgColor   string   `form:"sidebar_bg_color"`
	SidebarSubNames  []string `form:"sidebar_sub_name"`
	SidebarSubUrls   []string `form:"sidebar_sub_url"`
}

type SideBarJsonBodyRequest struct {
	IconID          uint     `json:"icon_id"`
	SidebarName     string   `json:"sidebar_name"`
	SidebarSubNames []string `json:"sidebar_sub_name"`
	SidebarSubUrls  []string `json:"sidebar_sub_url"`
}

type SidebarChildMainQueryRow struct {
	ID               uint                      `gorm:"id"`
	SidebarName      string                    `gorm:"sidebar_name"`
	IconID           uint                      `gorm:"icon_id"`
	SidebarTextColor string                    `gorm:"sidebar_text_color"`
	SidebarBgColor   string                    `gorm:"sidebar_bg_color"`
	SidebarSub       []SidebarChildSubQueryRow `gorm:"sidebar_sub"`
}

type SidebarChildSubQueryRow struct {
	ID              uint   `gorm:"id"`
	SidebarMainID   uint   `gorm:"sidebar_main_id"`
	SidebarSubNames string `gorm:"sidebar_sub_name"`
	SidebarSubUrls  string `gorm:"sidebar_sub_url"`
}

type SideBarRespone struct {
	ID                uint                `json:"id"`
	SidebarName       string              `json:"sidebar_name"`
	SidebarIcon       string              `json:"sidebar_icon"`
	SidebarTextColor  string              `json:"sidebar_text_color"`
	SidebarBgColor    string              `json:"sidebar_bg_color"`
	SidebarSubRespone []SidebarSubRespone `json:"sidebar_subs"`
}

type SidebarSubRespone struct {
	ID              uint   `json:"id"`
	SidebarMainID   uint   `json:"sidebar_main_id"`
	SidebarSubNames string `json:"sidebar_sub_name"`
	SidebarSubUrls  string `json:"sidebar_sub_url"`
}
