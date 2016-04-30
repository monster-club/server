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

func GetOne(ctrl controller.CRUDController) func(c *gin.Context) {
	return func(c *gin.Context) {
		res, err := ctrl.Find(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Bad id"})
		} else {
			c.JSON(200, res)
		}
	}
}
