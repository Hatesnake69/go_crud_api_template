package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/", helloWorldHandler)
}
