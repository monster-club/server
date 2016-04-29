package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Pokemon struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Abilities  []int32       `json:"abilities" bson:"abilities"`
	EggGroups  []int32       `json:"eggGroups" bson:"egg_groups"`
	EvYield    []int32       `json:"evYield" bson:"ev_yield"`
	Types      []int32       `json:"types" bson:"types"`
	Evolutions []Evolution   `json:"evolutions" bson:"evolutions"`
	Moves      []LearnSet    `json:"moves" bson:"moves"`
	CatchRate  int32         `json:"catchRate" bson:"catch_rate"`
	EggCycles  int32         `json:"eggCycles" bson:"egg_cycles"`
	Happiness  int32         `json:"happiness" bson:"happiness"`
	Exp        int32         `json:"exp" bson:"exp"`
	GrowthRate string        `json:"growthRate" bson:"growth_rate"`
	Height     float64       `json:"height" bson:"height"`
	Weight     float64       `json:"weight" bson:"weight"`
	Ratio      float64       `json:"ratio" bson:"ratio"`
	DexNum     int32         `json:"dexNum" bson:"dex_num"`
	Stats      Stat          `json:"stats" bson:"stats"`
}

func (p *Pokemon) Valid() bool {
	abilitiesLen := len(p.Abilities)
	eggGroupsLen := len(p.EggGroups)
	typesLen := len(p.Types)
	return (p.Name != "" &&
		(abilitiesLen >= 1 && abilitiesLen <= 3) &&
		(eggGroupsLen >= 1 && eggGroupsLen <= 2) &&
		(typesLen >= 1 && typesLen <= 2) &&
		len(p.Moves) != 0 &&
		p.CatchRate > 0 &&
		p.EggCycles > 0 &&
		p.Exp > 0 &&
		p.GrowthRate != "" &&
		p.Height > 0.0 &&
		p.Weight > 0.0 &&
		(p.Ratio >= 0.0 && p.Ratio <= 100.0) &&
		p.DexNum > 0 &&
		p.Stats.valid())
}
