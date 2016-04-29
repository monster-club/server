package model

type Move struct {
	Learn string `json:"learn" bson:"learn"`
	Level int32  `json:"level,omitempty" bson:"level,omitempty"`
	Num   int32  `json:"num" bson:"num"`
}
