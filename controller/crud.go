package controller

/*
import (
	"github.com/pokemon-club/server/model"
)
*/

type CRUDController interface {
	All() interface{}
	Find(id string, m interface{}) (interface{}, error)
	Insert(m interface{}) (interface{}, error)
	Update(id string, m interface{}) (interface{}, error)
	Delete(id string) error
}
