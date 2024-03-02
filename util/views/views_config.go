package util

import (
	"os"
	"time"
)

type config struct {
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Set() map[string]interface{} {
	return map[string]interface{}{
		"app_title":     os.Getenv("APP_NAME_WEB"),
		"app_name_html": os.Getenv("APP_NAME_HTML"),
		"base_url":      os.Getenv("APP_URL"),
		"script_time":   time.Now().Format("15:04:05"),
	}
}
