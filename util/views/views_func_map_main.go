package util

import (
	"github.com/gofiber/template/html/v2"
)

func FuncMap(engine *html.Engine) *html.Engine {
	config := NewConfig()
	engine.AddFunc("config", config.Set)

	sidebars := NewSidebar()
	engine.AddFunc("sidebars", sidebars.renderAll)

	icon := NewMasterIcon()
	engine.AddFunc("iconSelete2", icon.renderAll)
	return engine
}
