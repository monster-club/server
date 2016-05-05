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
	found, ok := res.(*model.Pokemon)
	if !ok {
		t.Error("Interface type could not be convereted to a Pokemon struct")
	}
	if found.Name != "Charmander" {
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

func TestInsertPokemonReturnsAnErrorForInvalidData(t *testing.T) {
	db, _, cont := standardInsertSetup()
	defer db.DropDatabase()
	newPkm := pokemonFactory()
	newPkm.DexNum = 0
	_, err := cont.Insert(&newPkm)
	if newPkm.Valid() == true {
		t.Error("Pokemon should be invalid")
	}
	if err == nil {
		t.Error("An error should have been raised for invalid data")
	}
}

func TestInsertPokemon(t *testing.T) {
	db, cont := standardSetup()
	defer db.DropDatabase()
	charmander := pokemonFactory()
	ret, err := cont.Insert(&charmander)
	if err != nil {
		t.Error("There was a database error trying to insert.", err)
	}
	pkm, ok := ret.(*model.Pokemon)
	if !ok {
		t.Error("Failed to convert interface to Pokemon struct.", ok)
	}
	// TODO: make it so that an id is entered on creation
	if bson.ObjectId.Hex(pkm.ID) != "" {
		t.Error("The inserted Pokemon should have an ID.")
	}
}

func TestUpdatePokemon(t *testing.T) {
	db, pkm, cont := standardInsertSetup()
	var m model.Pokemon
	defer db.DropDatabase()
	newPkm := pokemonFactory()
	newPkm.Name = "Squirtle"
	_, err := cont.Update(bson.ObjectId.Hex(pkm.ID), &newPkm)
	if err != nil {
		t.Error("There should not have been a database error:", err)
	}
	res, findErr := cont.Find(bson.ObjectId.Hex(pkm.ID), &m)
	found, _ := res.(*model.Pokemon)
	if findErr != nil {
		t.Error("Database error finding updated pokemon")
	}
	if found.Name != "Squirtle" {
		t.Error("Document was not updated")
	}
	if found.DexNum == 0 {
		t.Error("Document was mangled on update")
	}
}

func TestUpdateFailsWithBadId(t *testing.T) {
	db, _, cont := standardInsertSetup()
	defer db.DropDatabase()
	newPkm := pokemonFactory()
	newPkm.Name = "Bulbasaur"
	_, err := cont.Update("lol totally a hex", &newPkm)
	if err == nil {
		t.Error("An error should have been raised for an invalid hex")
	}
}

func TestUpdateFailsWithBadData(t *testing.T) {
	db, pkm, cont := standardInsertSetup()
	defer db.DropDatabase()
	newPkm := pokemonFactory()
	newPkm.DexNum = 0
	_, err := cont.Update(bson.ObjectId.Hex(pkm.ID), &newPkm)
	if newPkm.Valid() == true {
		t.Error("Pokemon should be invalid")
	}
	if err == nil {
		t.Error("An error should have been raised for invalid data")
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

func TestDeletePokemonReturnsAnErrorForBadIds(t *testing.T) {
	db, _, cont := standardInsertSetup()
	defer db.DropDatabase()
	err := cont.Delete("abc123")
	if err == nil {
		t.Error("An error should have been returned when a invalid hex was provided")
	}
}
