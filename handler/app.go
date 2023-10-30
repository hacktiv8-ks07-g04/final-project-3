package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hacktiv8-ks07-g04/final-project-3/config"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/database"
)

func StartApp() {
	config.LoadAppConfig()

	database.ConnectDB()

	r := gin.Default()
	ping := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
	r.GET("/ping", ping)
	r.Run(":8080")
}
