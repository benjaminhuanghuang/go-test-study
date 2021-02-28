package main

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.PingController.Ping)

	router.Run(":8080")
}
