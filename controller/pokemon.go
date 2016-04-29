package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
)

type Pokemon struct {
	C *mgo.Collection
}

// All is an exported method that returns all documents for a pokemon as
// Pokemon structs.
func (p *Pokemon) All() []model.Pokemon {
	var results []model.Pokemon
	_ = p.C.Find(nil).All(&results)
	return results
}

// NewPokemon is a factory function that creates a new Pokemon controller
// struct, initialized with the default collection that it is supposed to use.
// NewPokemon should be preferred over creating a controller with a struct
// literal like:
// controller := Pokemon{C: collection}
func NewPokemon(db *mgo.Database) *Pokemon {
	return &Pokemon{C: db.C("Pokemon")}
}
