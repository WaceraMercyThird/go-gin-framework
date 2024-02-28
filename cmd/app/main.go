package main

import (
	"github.com/gin-gonic/gin"
)

// @title Documenting API (Report)
// @version 1
// @Description Sample description

// @contact.name MWacera
// @contact.url https:github.com/WaceraMercyThird
// @contact.email waceraaamercy@gmail.com
// @


// @host localhost:8080
// @BasePath /api/v1
// @
// @
// @

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) { 
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}
