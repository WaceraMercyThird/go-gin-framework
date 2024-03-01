package main

import (
	"log"

	_ "github.com/WaceraMercyThird/go_gin_project_swagger/cmd/app/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API (Report)
// @version 1



// @host localhost:8080
// @BasePath /api/v1
// @
// @
// @

func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run()

	if err != nil {
		log.Fatal(err)
	}

}

func routermain() {
	router := gin.Default()

	router.GET("/swagger", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}
