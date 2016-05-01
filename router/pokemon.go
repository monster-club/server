package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon-club/server/controller"
	"github.com/pokemon-club/server/model"
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

func Create(ctrl controller.CRUDController) func(c *gin.Context) {
	var json model.Pokemon
	return func(c *gin.Context) {
		if c.BindJSON(&json) == nil {
			res, err := ctrl.Insert(json)
			if err == nil {
				return c.JSON(200, res)
			} else {
				return c.JSON(400, gin.H{"error": err})
			}
		} else {
			return c.JSON(500, gin.H{"error": "Couldn't convert data"})
		}
	}
}
