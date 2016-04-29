package model

import (
	"testing"
)

// Creates a valid pokemon and returns it.
func pokemonFactory() Pokemon {
	return Pokemon{
		Name: "Charmander",
		Abilities: []int32{1, 2},
		EggGroups: []int32{1},
		Types: []int32{1},
		Moves: []Move{Move{Learn: "a", Level: 1, Num: 1}},
		CatchRate: 1,
		EggCycles: 1,
		Exp: 1,
		GrowthRate: "medium_slow",
		Height: 1.0,
		Weight: 1.0,
		Ratio: 87.5,
		DexNum: 1,
	}
}

// Assert that a Pokemon struct we expect to be valid is.
func TestValidity(t *testing.T) {
	p := pokemonFactory()
	if p.valid() != true {
		t.Error("Expected pokemon to be valid.")
	}
}

func TestNameMustExist(t *testing.T) {
	p := pokemonFactory()
	p.Name = ""
	if p.valid() == true {
		t.Error("Pokemon should not be valid without a name.")
	}
}

func TestAbilitiesMustBeAtLeastOne(t *testing.T) {
	p := pokemonFactory()
	p.Abilities = []int32{}
	if p.valid() == true {
		t.Error("Pokemon should not be valid with 0 abilities.")
	}
}

func TestAbilitiesMustNotBeGreaterThanThree(t *testing.T) {
	p := pokemonFactory()
	p.Abilities = []int32{1, 2, 3, 4}
	if p.valid() == true {
		t.Error("Pokemon should not be valid with more than 3 abilities.")
	}
}

func TestEggGroupsMustBeAtLeastOne(t *testing.T) {
	p := pokemonFactory()
	p.EggGroups = []int32{}
	if p.valid() == true {
		t.Error("Pokemon should not be valid with 0 egg groups.")
	}
}

func TestEggGroupsMustNotBeGreaterThanTwo(t *testing.T) {
	p := pokemonFactory()
	p.EggGroups = []int32{1, 2, 3}
	if p.valid() == true {
		t.Error("Pokemon should not be valid with more than 2 egg groups.")
	}
}

func TestMovesMustBeAtLeastOne(t *testing.T) {
	p := pokemonFactory()
	p.Moves = []Move{}
	if p.valid() == true {
		t.Error("Pokemon should not be valid if they have no moves.")
	}
}

func TestGenderRatioMustNotBeNegative(t *testing.T) {
	p := pokemonFactory()
	p.Ratio = -1.0
	if p.valid() == true {
		t.Error("Pokemon can't have a negative gender ratio.")
	}
}

func TestGenderRatioMustNotBeGreaterThanOneHundred(t *testing.T) {
	p := pokemonFactory()
	p.Ratio = 101.0
	if p.valid() == true {
		t.Error("Pokemon can't have a gender ratio over 100 percent.")
	}
}
