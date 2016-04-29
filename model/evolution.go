package model

type Evolution struct {
	Method string `json:"method" bson:"method"`
	To     int32  `json:"to" bson:"to"`
}
