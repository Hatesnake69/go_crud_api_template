package main

import (
	"go_crud_api_template/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	endpoints.SetupRoutes(router)
	router.Run(":8080")
}
