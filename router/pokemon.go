package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon-club/server/controller"
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2/bson"
)

func GetAll(ctrl controller.CRUDController) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, ctrl.All())
	}
}

func GetOne(ctrl controller.CRUDController, m model.Document) func(c *gin.Context) {
	return func(c *gin.Context) {
		res, err := ctrl.Find(c.Param("id"), m)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, res)
		}
	}
}

func Create(ctrl controller.CRUDController, m model.Document) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.BindJSON(&m) == nil {
			res, err := ctrl.Insert(m)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
			} else {
				c.JSON(201, res)
			}
		} else {
			c.JSON(500, gin.H{"error": "Couldn't convert data"})
		}
	}
}

func Update(ctrl controller.CRUDController) func(c *gin.Context) {
	var json bson.M
	return func(c *gin.Context) {
		if c.BindJSON(&json) == nil {
			_, err := ctrl.Update(c.Param("id"), json)
			if err != nil {
				c.JSON(400, gin.H{"error": "Failed to insert data"})
			} else {
				c.AbortWithStatus(204)
			}
		} else {
			c.JSON(500, gin.H{"error": "Couldn't convert data"})
		}
	}
}

func Delete(ctrl controller.CRUDController) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := ctrl.Delete(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err})
		} else {
			c.AbortWithStatus(204)
		}
	}
}
