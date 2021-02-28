package controllers

import (
	"net/http"

	"services"

	"github.com/gin-gonic/gin"
)

var (
	PingController = pingController{}
)

type pingController struct{}

func (controller pingController) Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, result)
}
