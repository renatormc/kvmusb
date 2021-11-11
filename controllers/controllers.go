package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/kvmusb/services"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ListRunningVms(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SaveDevices(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SnapshotDevices(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func NewDevices(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AttachUsb(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DetachUsb(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AttachDisk(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DetachDisk(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ListUsb(c *gin.Context) {
	usbs, err := services.ListUsbs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, usbs)
}
