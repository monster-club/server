package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon-club/server/controller"
	"github.com/pokemon-club/server/model"
	"github.com/pokemon-club/server/router"
	"gopkg.in/mgo.v2"
)

func main() {
	r := gin.Default()
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	db := sess.DB("pokemon")
	pkmCtrl := controller.NewPokemon(db)

	var m model.Pokemon
	r.GET("/pokemon", router.GetAll(pkmCtrl))
	r.POST("/pokemon", router.Create(pkmCtrl, &m))
	r.GET("/pokemon/:id", router.GetOne(pkmCtrl, &m))
	r.PUT("/pokemon/:id", router.Update(pkmCtrl))
	r.DELETE("/pokemon/:id", router.Delete(pkmCtrl))

	r.Run()
}
