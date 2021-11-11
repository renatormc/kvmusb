package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/renatormc/kvmusb/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		main.GET("/test", controllers.Test)
		main.GET("/list-running-vms", controllers.ListRunningVms)
		main.GET("/save-devices", controllers.SaveDevices)
		main.GET("/snapshot-devices", controllers.SnapshotDevices)
		main.GET("/new-devices", controllers.NewDevices)
		main.GET("/attach-usb", controllers.AttachUsb)
		main.GET("/attach-disk", controllers.AttachDisk)
		main.GET("/dettach-disk", controllers.DetachDisk)
		main.GET("/dettach-usb", controllers.DetachUsb)
	}

	return router
}
