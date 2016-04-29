package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon-club/server/controller"
	"gopkg.in/mgo.v2"
)

func main() {
	r := gin.Default()
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	r.GET("/pokemon", func(c *gin.Context) {
		cont := controller.Pokemon{sess.DB("pokemon").C("pokemon")}
		c.JSON(200, cont.All())
	})

	r.Run()
}
