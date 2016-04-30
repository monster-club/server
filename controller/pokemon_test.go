package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func testInsert(db *mgo.Database) model.Pokemon {
	var charmander model.Pokemon
	charmander.Name = "Charmander"
	// Force an ID onto the record before hand
	charmander.ID = bson.NewObjectId()
	coll := db.C("pokemon")
	coll.Insert(&charmander)
	return charmander
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
	defer db.DropDatabase()
	_ = testInsert(db)
	cont := NewPokemon(db)
	all := cont.All()
	pkm, ok := all.([]model.Pokemon)
	if !ok {
		t.Error("Interface could not be convereted to a Pokemon model")
	}
	if len(pkm) != 1 {
		t.Error("The controller did not return all records.")
	}
	if pkm[0].Name != "Charmander" {
		t.Error("The controller did not return data intact.")
	}
}

func TestFindPokemon(t *testing.T) {
	db := mangoSetup()
	defer db.DropDatabase()
	pkm := testInsert(db)
	cont := NewPokemon(db)
	res, err := cont.find(bson.ObjectId.Hex(pkm.ID))
	if err != nil {
		t.Error("An error occured retrieving the Pokemon:", err)
	}
	if res.Name != "" && res.Name != pkm.Name {
		t.Error("The Id found the wrong entry.")
	}
}

func TestFindPokemonBadId(t *testing.T) {
	db := mangoSetup()
	defer db.DropDatabase()
	_ = testInsert(db)
	cont := NewPokemon(db)
	_, err := cont.find("lolol not a hex id")
	if err == nil {
		t.Error("An invalid hex id should have raised an error")
	}
}

func TestFindPokemonDatabaseProblem(t *testing.T) {
	db := mangoSetup()
	pkm := testInsert(db)
	cont := NewPokemon(db)
	// Drop the database early, so no records exist
	db.DropDatabase()
	_, err := cont.find(bson.ObjectId.Hex(pkm.ID))
	if err == nil {
		t.Error("An empty database should have raised an error")
	}
}
