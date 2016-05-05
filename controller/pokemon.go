package controller

import (
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
)

// The Pokemon controller struct is an impelmenter of the CRUDController
// interface. It is intended to be given to a router and handle database
// interactions.
type Pokemon struct {
	GenericCRUD
	C *mgo.Collection
}

// All is part of the CRUDController interface.
// All returns all documents from the controller's collection
func (p *Pokemon) All() interface{} {
	var results []model.Pokemon
	_ = p.C.Find(nil).All(&results)
	return results
}
