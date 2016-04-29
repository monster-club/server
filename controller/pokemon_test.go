package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
	"reflect"
	"testing"
)

func mangoSetup() *mgo.Database {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	return sess.DB("pokemonTest")
}

func TestPokemonNew(t *testing.T) {
	db := mangoSetup()
	cont := NewPokemon(db)
	reflection := reflect.TypeOf(cont)
	if reflection.String() != "*controller.Pokemon" {
		t.Error("The factory function created the wrong controller. Created:", reflection)
	}
}

func TestAllPokemon(t *testing.T) {
	db := mangoSetup()
	coll := db.C("pokemon")
	coll.Insert(&model.Pokemon{Name: "Charmander"})
	cont := Pokemon{C: coll}
	all := cont.All()
	if len(all) != 1 {
		t.Error("The controller did not return all records.")
	}
	if all[0].Name != "Charmander" {
		t.Error("The controller did not return data intact.")
	}
	db.DropDatabase()
}
