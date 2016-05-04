package model

// A document is an interface representing a root level document, that is, an
// object that goes directly into a collection, as opposed to an sub object in
// another document. Every document interface has to implement a Valid() method
// which is used to ensure that documents being inserted or updated have the
// right keys and values. And that the values are in expected ranges.
type Document interface {
	Valid() bool
}
