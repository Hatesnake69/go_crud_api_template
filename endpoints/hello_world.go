package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Hello World
// @Description Get a hello message
// @ID get-hello-world
// @Produce  json
// @Success 200 {object} string "OK"
// @Router /hello [get]
func HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/hello", HelloWorldHandler)

	swaggerURL := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
}
