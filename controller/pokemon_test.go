package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"testing"
)

func TestPokemonNew(t *testing.T) {
	_, cont := standardSetup()
	reflection := reflect.TypeOf(cont)
	if reflection.String() != "*controller.Pokemon" {
		t.Error("The factory function created the wrong controller. Created:", reflection)
	}
}

func TestAllPokemon(t *testing.T) {
	db, _, cont := standardInsertSetup()
	defer db.DropDatabase()
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
	db, pkm, cont := standardInsertSetup()
	var m model.Pokemon
	defer db.DropDatabase()
	res, err := cont.Find(bson.ObjectId.Hex(pkm.ID), &m)
	if err != nil {
		t.Error("An error occured retrieving the Pokemon:", err)
	}
	found, ok := res.(bson.M)
	if !ok {
		t.Error("Interface type could not be convereted to a Pokemon struct")
	}
	if found["name"] != "Charmander" {
		t.Error("Data did not come out of transformation intact.")
	}
}

func TestFindPokemonBadId(t *testing.T) {
	db, _, cont := standardInsertSetup()
	var m model.Pokemon
	defer db.DropDatabase()
	_, err := cont.Find("lolol not a hex id", &m)
	if err == nil {
		t.Error("An invalid hex id should have raised an error")
	}
}

func TestFindPokemonDatabaseProblem(t *testing.T) {
	db, pkm, cont := standardInsertSetup()
	// Drop the database early, so no records exist
	var m model.Pokemon
	db.DropDatabase()
	_, err := cont.Find(bson.ObjectId.Hex(pkm.ID), &m)
	if err == nil {
		t.Error("An empty database should have raised an error")
	}
}

func TestInsertPokemonThrowsAnErrorForInvalidStruct(t *testing.T) {
	db, charmander, cont := standardInsertSetup()
	defer db.DropDatabase()
	_, err := cont.Insert(&charmander)
	if err == nil {
		t.Error("An invalid struct should raise an error.", err)
	}
}

func TestInsertPokemon(t *testing.T) {
	db, cont := standardSetup()
	defer db.DropDatabase()
	charmander := pokemonFactory()
	ret, err := cont.Insert(charmander)
	if err != nil {
		t.Error("There was a database error trying to insert.", err)
	}
	pkm, ok := ret.(model.Pokemon)
	if !ok {
		t.Error("Failed to convert interface to Pokemon struct.", ok)
	}
	if bson.ObjectId.Hex(pkm.ID) == "" {
		t.Error("The inserted Pokemon should have an ID.")
	}
}

func TestInsertPokemonBadInterface(t *testing.T) {
	db, cont := standardSetup()
	defer db.DropDatabase()
	wrongThing := Pokemon{}
	_, err := cont.Insert(wrongThing)
	if err == nil {
		t.Error("Should have failed to convert to a Pokemon struct.", err)
	}
}

func TestUpdatePokemon(t *testing.T) {
	db, pkm, cont := standardInsertSetup()
	var m model.Pokemon
	defer db.DropDatabase()
	_, err := cont.Update(bson.ObjectId.Hex(pkm.ID), bson.M{"name": "Squirtle"})
	if err != nil {
		t.Error("There should not have been a database error:", err)
	}
	res, findErr := cont.Find(bson.ObjectId.Hex(pkm.ID), &m)
	found, _ := res.(bson.M)
	if findErr != nil {
		t.Error("Database error finding updated pokemon")
	}
	if found["name"] != "Squirtle" {
		t.Error("Document was not updated")
	}
	if found["dex_num"] == 0 {
		t.Error("Document was mangled on update")
	}
}

func TestDeletePokemon(t *testing.T) {
	db, charmander, cont := standardInsertSetup()
	defer db.DropDatabase()
	err := cont.Delete(bson.ObjectId.Hex(charmander.ID))
	if err != nil {
		t.Error("An error occurred trying to delete the pokemon", err)
	}
}
