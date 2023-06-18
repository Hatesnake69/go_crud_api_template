package main

import (
	"go_crud_api_template/docs"
	"go_crud_api_template/endpoints"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.Title = "Your API Title"
	docs.SwaggerInfo.Description = "Your API Description"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	router := gin.Default()

	// Register the Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Use a different path for your hello world endpoint
	router.GET("/hello", endpoints.HelloWorldHandler)

	router.Run(":8080")
}
