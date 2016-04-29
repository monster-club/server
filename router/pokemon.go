package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon-club/server/controller"
)

func GetAll(ctrl controller.CRUDController) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, ctrl.All())
	}
}
