package handler

import (
	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()
	ping := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
	r.GET("/ping", ping)
	r.Run(":8080")
}
