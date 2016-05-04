package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func mangoSetup() *mgo.Database {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	return sess.DB("pokemonTest")
}

// Returns the three most common things needed in the controller integration test, a database
// a struct that has been inserted into that database, and a controller reference.
func standardSetup() (*mgo.Database, *Pokemon) {
	db := mangoSetup()
	cont := NewPokemon(db)
	return db, cont
}

func standardInsertSetup() (*mgo.Database, model.Pokemon, *Pokemon) {
	db, cont := standardSetup()
	pkm := testInsertValid(db)
	return db, pkm, cont
}

func testInsertValid(db *mgo.Database) model.Pokemon {
	charmander := pokemonFactory()
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
