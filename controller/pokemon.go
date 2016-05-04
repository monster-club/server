package controller

import (
	"errors"
	"github.com/pokemon-club/server/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// The Pokemon controller struct is an impelmenter of the CRUDController
// interface. It is intended to be given to a router and handle database
// interactions.
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
		err := p.C.FindId(bson.ObjectIdHex(id)).One(&result)
		return result, err
	}
	return model.Pokemon{}, errors.New("Invalid id provided")
}

// insert will create a new document in the collection for the given Pokemon
// struct, provided the struct is valid. If the struct is invalid, or if there
// is a database error trying to insert the data, an error will be returned.
func (p *Pokemon) insert(m *model.Pokemon) (model.Pokemon, error) {
	if m.Valid() == false {
		return model.Pokemon{}, errors.New("Invalid data for creation.")
	}
	m.ID = bson.NewObjectId()
	err := p.C.Insert(m)
	return *m, err
}

// update will change toe provided keys in the "m" variable. It will also
// return an error, if there is an error on the database level during the
// update, or if the id provided is not a valid ObjectId hex.
func (p *Pokemon) update(id string, m bson.M) (bson.M, error) {
	if bson.IsObjectIdHex(id) == true {
		err := p.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": m})
		return m, err
	}
	return bson.M{}, errors.New("Invalid id provided")
}

// delete takes in a Object Id hex string for a document in the collection.
// if it is successful it will return nil, otherwise it will return a database
// error. If an invalid hex string was given, it will return an error
// indicating that.
func (p *Pokemon) delete(id string) error {
	if bson.IsObjectIdHex(id) == true {
		err := p.C.RemoveId(bson.ObjectIdHex(id))
		return err
	}
	return errors.New("Invalid id provided")
}

// All is an exported method that returns all documents from the pokemon
// collection as an interface{}. This will need to be cast to a model.Pokemon.
// All is part of the CRUDController interface.
func (p *Pokemon) All() interface{} {
	return p.all()
}

// Find is a part of the CRUDController interface. It is a passthru for the
// local "find" method.
func (p *Pokemon) Find(id string) (interface{}, error) {
	return p.find(id)
}

// Insert is a part of the CRUDController interface. It is a passthru for the
// local "insert" method.
func (p *Pokemon) Insert(m interface{}) (interface{}, error) {
	pkm, ok := m.(model.Pokemon)
	if !ok {
		return model.Pokemon{}, errors.New("Failed to convert interface.")
	}
	return p.insert(&pkm)
}

// Update is a part of the CRUDController interface. It is a passthru to the
// unexported .update call.
func (p *Pokemon) Update(id string, m interface{}) (interface{}, error) {
	pkm, ok := m.(bson.M)
	if !ok {
		return bson.M{}, errors.New("Bad interface")
	}
	return p.update(id, pkm)
}

// Delete is a part of the CRUDController interface. it is a passthru to the
// undexported .delete call.
func (p *Pokemon) Delete(id string) error {
	return p.delete(id)
}
