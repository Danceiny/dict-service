package api

import (
    "github.com/Danceiny/dict-service/controller"
    "github.com/gin-gonic/gin"
)

var dictHandler = &controller.DictHandler{}

var Server = gin.Default()

func init() {
    Server.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    Server.GET("/api/dict/area/:id", dictHandler.GetArea)
}
