package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
)

type Pokemon struct {
	C *mgo.Collection
}

// all returns all documents from the pokemon collection as model.Pokemon
// structs.
func (p *Pokemon) all() []model.Pokemon {
	var results []model.Pokemon
	_ = p.C.Find(nil).All(&results)
	return results
}

// All is an exported method that returns all documents from the pokemon
// collection as an interface{}. This will need to be cast to a model.Pokemon.
// All is part of the CRUDController interface.
func (p *Pokemon) All() interface{} {
	return p.all()
}
