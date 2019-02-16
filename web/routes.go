package web

import (
	. "github.com/Danceiny/dict-service/controller"
	"github.com/gin-gonic/gin"
)

var (
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	Server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Server.GET("/api/dict/area/:id", DictControllerCpt.GetArea)
}
