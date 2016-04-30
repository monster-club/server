package controller

import (
	"errors"
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// find will retrieve a single document from the pokemon collection, and return
// it as a Pokemon struct. If the provided id is invalid, or an error occurs
// in the database during retrieval, an error will be returned as well.
func (p *Pokemon) find(id string) (model.Pokemon, error) {
	if bson.IsObjectIdHex(id) == true {
		var result model.Pokemon
		err := p.C.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
		return result, err
	} else {
		return model.Pokemon{}, errors.New("Invalid id provided")
	}
}

// Find is a part of the CRUDController interface. It is a passthru for the
// local "find" method.
func (p *Pokemon) Find(id string) (interface{}, error) {
	return p.find(id)
}

// All is an exported method that returns all documents from the pokemon
// collection as an interface{}. This will need to be cast to a model.Pokemon.
// All is part of the CRUDController interface.
func (p *Pokemon) All() interface{} {
	return p.all()
}
