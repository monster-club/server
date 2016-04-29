package model

type Stat struct {
	Hp             int32 `json:"hp" bson:"hp"`
	Attack         int32 `json:"attack" bson:"attack"`
	Defense        int32 `json:"defense" bson:"defense"`
	SpecialAttack  int32 `json:"special_attack" bson:"special_attack"`
	SpecialDefense int32 `json:"special_defense" bson:"special_defense"`
	Speed          int32 `json:"speed" bson:"speed"`
}
