package controller

type CRUDController interface {
	All() interface{}
	Find(id string) (interface{}, error)
	Insert(m interface{}) (interface{}, error)
	Update(id string, m interface{}) (interface{}, error)
	Delete(id string) error
}
