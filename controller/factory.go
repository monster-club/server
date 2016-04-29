package controller

import (
	"gopkg.in/mgo.v2"
)

// NewPokemon is a factory function that creates a new Pokemon controller
// struct, initialized with the default collection that it is supposed to use.
// NewPokemon should be preferred over creating a controller with a struct
// literal like:
// controller := Pokemon{C: collection}
func NewPokemon(db *mgo.Database) *Pokemon {
	return &Pokemon{C: db.C("pokemon")}
}
