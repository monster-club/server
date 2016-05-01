package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"testing"
)

func mangoSetup() (*mgo.Database, error) {
	sess, err := mgo.Dial("127.0.0.1")
	return sess.DB("pokemonTest"), err
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

// Returns a model.Pokemon struct that will return true from a valid() call
func pokemonFactory() model.Pokemon {
	return model.Pokemon{
		Name:       "Charmander",
		Abilities:  []int32{1, 2},
		EggGroups:  []int32{1},
		Types:      []int32{1},
		Moves:      []model.LearnSet{model.LearnSet{Learn: "a", Level: 1, Num: 1}},
		CatchRate:  1,
		EggCycles:  1,
		Exp:        1,
		GrowthRate: "medium_slow",
		Height:     1.0,
		Weight:     1.0,
		Ratio:      87.5,
		DexNum:     1,
		Stats:      model.Stat{1, 1, 1, 1, 1, 1},
	}
}

func TestPokemonNew(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		cont := NewPokemon(db)
		reflection := reflect.TypeOf(cont)
		if reflection.String() != "*controller.Pokemon" {
			t.Error("The factory function created the wrong controller. Created:", reflection)
		}
	} else {
		t.Error("Unable to create database connection:", err)
	}
}

func TestAllPokemon(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
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
	} else {
		t.Error("Unable to create database connection:", err)
	}
}

func TestFindPokemon(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		defer db.DropDatabase()
		pkm := testInsert(db)
		cont := NewPokemon(db)
		res, err := cont.Find(bson.ObjectId.Hex(pkm.ID))
		if err != nil {
			t.Error("An error occured retrieving the Pokemon:", err)
		}
		found, ok := res.(model.Pokemon)
		if !ok {
			t.Error("Interface type could not be convereted to a Pokemon struct")
		}
		if found.Name != "Charmander" {
			t.Error("Data did not come out of transformation intact.")
		}
	}
}

func TestFindPokemonBadId(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		defer db.DropDatabase()
		_ = testInsert(db)
		cont := NewPokemon(db)
		_, err := cont.find("lolol not a hex id")
		if err == nil {
			t.Error("An invalid hex id should have raised an error")
		}
	} else {
		t.Error("Unable to create database connection:", err)
	}
}

func TestFindPokemonDatabaseProblem(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		pkm := testInsert(db)
		cont := NewPokemon(db)
		// Drop the database early, so no records exist
		db.DropDatabase()
		_, err := cont.find(bson.ObjectId.Hex(pkm.ID))
		if err == nil {
			t.Error("An empty database should have raised an error")
		}
	} else {
		t.Error("Unable to create database connection:", err)
	}
}

func TestInsertPokemonThrowsAnErrorForInvalidStruct(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		defer db.DropDatabase()
		cont := NewPokemon(db)
		charmander := model.Pokemon{Name: "Charmander"}
		_, err := cont.Insert(&charmander)
		if err == nil {
			t.Error("An invalid struct should raise an error.", err)
		}
	} else {
		t.Error("Unable to create database connection:", err)
	}
}

func TestInsertPokemon(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		defer db.DropDatabase()
		cont := NewPokemon(db)
		charmander := pokemonFactory()
		ret, err := cont.Insert(charmander)
		if err != nil {
			t.Error("There was a database error trying to insert.", err)
		}
		pkm, ok := ret.(model.Pokemon)
		if !ok {
			t.Error("Failed to convert interface to Pokemon struct.", ok)
		}
		hex := bson.ObjectId.Hex(pkm.ID)
		if hex == "" {
			t.Error("The inserted Pokemon should have an ID.", hex)
		}
	}
}

func TestInsertPokemonBadInterface(t *testing.T) {
	db, err := mangoSetup()
	if err == nil {
		defer db.DropDatabase()
		cont := NewPokemon(db)
		wrongThing := Pokemon{}
		_, err := cont.Insert(wrongThing)
		if err == nil {
			t.Error("Should have failed to convert to a Pokemon struct.", err)
		}
	}
}
