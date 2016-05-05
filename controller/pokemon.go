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

// All is part of the CRUDController interface.
// All returns all documents from the controller's collection
func (p *Pokemon) All() interface{} {
	var results []model.Pokemon
	_ = p.C.Find(nil).All(&results)
	return results
}

// Find is a part of the CRUDController interface.
// Find will retrieve a single document from the pokemon collection, and return
// it as an interface. If the provided id is invalid, or an error occurs
// in the database during retrieval, an error will be returned as well.
func (p *Pokemon) Find(id string, m interface{}) (interface{}, error) {
	if bson.IsObjectIdHex(id) == true {
		err := p.C.FindId(bson.ObjectIdHex(id)).One(&m)
		return m, err
	}
	return model.Pokemon{}, errors.New("Invalid id provided")
}

// Insert is a part of the CRUDController interface.
// Insert will create a new document in the collection for the given document
// provided the document is valid. If the document is invalid, or if there
// is a database error trying to insert the data, an error will be returned.
func (p *Pokemon) Insert(m model.Document) (model.Document, error) {
	if m.Valid() == false {
		return &model.Pokemon{}, errors.New("Invalid data for creation.")
	}
	err := p.C.Insert(m)
	return m, err
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
