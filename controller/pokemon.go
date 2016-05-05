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
func (p *Pokemon) Find(id string, m model.Document) (model.Document, error) {
	if bson.IsObjectIdHex(id) == true {
		err := p.C.FindId(bson.ObjectIdHex(id)).One(m)
		return m, err
	}
	return &model.Pokemon{}, errors.New("Invalid id provided")
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

// Update is a part of the CRUDController interface.
// Update will do a full update of the selected document. Before writing, the
// new document will check for validity. If the new document would be invalid,
// Update will return an error. An error will also be returned in the case where
// and invalid Object Id hex is provided. Finally, an error will also be
// returned if there is an error writing to the database.
func (p *Pokemon) Update(id string, m model.Document) (model.Document, error) {
	if bson.IsObjectIdHex(id) == true {
		if m.Valid() == true {
			err := p.C.UpdateId(bson.ObjectIdHex(id), m)
			return m, err
		}
		return &model.Pokemon{}, errors.New("Provided object is not valid")
	}
	return &model.Pokemon{}, errors.New("Invalid id provided")
}

// Delete is a part of the CRUDController interface.
// Delete takes in a Object Id hex string for a document in the collection.
// If it is successful it will return nil, otherwise it will return a database
// error. If an invalid hex string was given, it will return an error
// indicating that.
func (p *Pokemon) Delete(id string) error {
	if bson.IsObjectIdHex(id) == true {
		err := p.C.RemoveId(bson.ObjectIdHex(id))
		return err
	}
	return errors.New("Invalid id provided")
}
