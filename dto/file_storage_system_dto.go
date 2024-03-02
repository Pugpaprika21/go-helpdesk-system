package dto

type FileStorageSystemQueryRow struct {
	ID            uint   `gorm:"id"`
	FileName      string `gorm:"file_name"`
	FileSize      int64  `gorm:"file_size"`
	FileType      string `gorm:"file_type"`
	FileExtension string `gorm:"file_extension"`
	Content       string `gorm:"content"`
	RefID         uint   `gorm:"ref_id"`
	RefTable      string `gorm:"ref_table"`
	RefField      string `gorm:"ref_field"`
}

type FileStorageSystemRespone struct {
	ID            uint   `json:"id"`
	FileName      string `json:"fileName"`
	FileSize      int64  `json:"fileSize"`
	FileType      string `json:"fileType"`
	FileExtension string `json:"fileExtension"`
	Content       string `json:"content,omitempty"`
	RefID         uint   `json:"refId,omitempty"`
	RefTable      string `json:"refTable,omitempty"`
	RefField      string `json:"refField,omitempty"`
}
