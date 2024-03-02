package router

import (
	"github.com/Pugpaprika21/go-fiber/app/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	sidebarController := controller.NewSidebarController()
	masterIconController := controller.NewMasterIconController()
	masterRepairEquipmentCategoryController := controller.NewMasterRepairEquipmentCategoryController()
	repairController := controller.NewRepairController()

	webRouter := app.Group("/web")
	{
		webRouter.Get("/index", func(c *fiber.Ctx) error {
			return c.Render("pages/admin/index", fiber.Map{}, "layouts/admin/main")
		})

		pageMasterRouter := webRouter.Group("/master")
		pageMasterRouter.Get("/sidebar_main_create", sidebarController.SidebarMainCreate)
		pageMasterRouter.Get("/sidebar_menu_create", sidebarController.SidebarCreate)
		pageMasterRouter.Get("/icon_create", masterIconController.IconCreate)

		pageRouter := webRouter.Group("/page")
		pageRouter.Get("/repair_form_save", repairController.RepairFormSave)
		pageRouter.Get("/repair_list", repairController.RepairList)
		pageRouter.Get("/repair_form_edit", repairController.RepairFormEdit)

		pageMasterRepairRouter := pageRouter.Group("/master")
		pageMasterRepairRouter.Get("/equipment_category_manage", masterRepairEquipmentCategoryController.RepairEquipmentCategoryCreate)
	}

	apiRouter := app.Group("/api")
	{
		masterRouter := apiRouter.Group("/master")
		masterRouter.Get("/sidebar_all", sidebarController.SidebarAll)
		masterRouter.Post("/sidebar_main_save", sidebarController.SidebarMainSave)
		masterRouter.Post("/sidebar_menu_save", sidebarController.SidebarMenuSave)
		masterRouter.Get("/icon_all", masterIconController.IconAll)
		masterRouter.Post("/icon_save", masterIconController.IconSave)

		masterRouter.Post("/equipment_category_save", masterRepairEquipmentCategoryController.EquipmentCategorySave)
		masterRouter.Get("/get_equipment_category_list", masterRepairEquipmentCategoryController.GetEquipmentCategoryList)

		repairRouter := apiRouter.Group("/repair")
		repairRouter.Post("/repair_save", repairController.RepairSave)
		repairRouter.Get("/repair_attached_files", repairController.RepairAttachedFiles)
		repairRouter.Get("/repair_get_list", repairController.RepairGetList)
		repairRouter.Delete("/repair_delete/:repair_id", repairController.RepairDelete)
		repairRouter.Delete("/repair_delete_attached_files/:file_id/:file_name", repairController.RepairDeleteAttachedFile)
		repairRouter.Get("/repair_edit/:repair_id", repairController.RepairEdit)
		repairRouter.Put("/repair_edit_save/:repair_id", repairController.RepairEditSave)
	}
}
