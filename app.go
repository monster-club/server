package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {
	r := gin.Default()
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil { panic(err) }
	defer sess.Close()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H { "story": "cool", "person": "bro", })
	})

	r.Run()
}
