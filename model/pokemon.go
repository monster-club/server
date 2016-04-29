package model

// ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
type Pokemon struct {
	Name       string      `json:"name" bson:"name"`
	Abilities  []int32     `json:"abilities" bson:"abilities"`
	EggGroups  []int32     `json:"egg_groups" bson:"egg_groups"`
	EvYield    []int32     `json:"ev_yield" bson:"ev_yield"`
	Types      []int32     `json:"types" bson:"types"`
	Evolutions []Evolution `json:"evolutions" bson:"evolutions"`
	Moves      []Move      `json:"moves" bson:"moves"`
	CatchRate  int32       `json:"catch_rate" bson:"catch_rate"`
	EggCycles  int32       `json:"egg_cycles" bson:"egg_cycles"`
	Happiness  int32       `json:"happiness" bson:"happiness"`
	Exp        int32       `json:"exp" bson:"exp"`
	GrowthRate string      `json:"growth_rate" bson:"growth_rate"`
	Height     float64     `json:"height" bson:"height"`
	Weight     float64     `json:"weight" bson:"weight"`
	Ratio      float64     `json:"ratio" bson:"ratio"`
	DexNum     int32       `json:"dex_num" bson:"dex_num"`
	Stats      Stat        `json:"stats" bson:"stats"`
}

func (p *Pokemon) valid() bool {
	abilitiesLen := len(p.Abilities)
	eggGroupsLen := len(p.EggGroups)
	typesLen := len(p.Types)
	return (p.Name != "" &&
		(abilitiesLen >= 1 && abilitiesLen <= 3) &&
		(eggGroupsLen >= 1 && eggGroupsLen <= 2) &&
		(typesLen >= 1 && typesLen <= 2) &&
		len(p.Moves) != 0 &&
		p.CatchRate != 0 &&
		p.EggCycles != 0 &&
		p.Exp != 0 &&
		p.GrowthRate != ""&&
		p.Height != 0.0 &&
		p.Weight != 0.0 &&
		(p.Ratio >= 0.0 && p.Ratio <= 100.0) &&
		p.DexNum != 0 &&
		p.Stats.valid())
}
