package model

import (
	"testing"
)

func statFactory() Stat {
	return Stat{
		Hp: 1,
		Attack: 1,
		Defense: 1,
		SpecialAttack: 1,
		SpecialDefense: 1,
		Speed: 1,
	}
}

func TestStatValidity(t *testing.T) {
	s := statFactory()
	if s.valid() == false {
		t.Error("Expected stats to be valid.")
	}
}

func TestHp(t *testing.T) {
	s := statFactory()
	s.Hp = 0
	if s.valid() == true {
		t.Error("An HP stat of 0 should not be valid.")
	}
}

func TestAttack(t *testing.T) {
	s := statFactory()
	s.Attack = 0
	if s.valid() == true {
		t.Error("An attack stat of 0 should not be valid.")
	}
}

func TestDefense(t *testing.T) {
	s := statFactory()
	s.Defense = 0
	if s.valid() == true {
		t.Error("An defense stat of 0 should not be valid.")
	}
}

func TestSpecialAttack(t *testing.T) {
	s := statFactory()
	s.SpecialAttack = 0
	if s.valid() == true {
		t.Error("A special attack stat of 0 should not be valid.")
	}
}

func TestSpecialDefense(t *testing.T) {
	s := statFactory()
	s.SpecialDefense = 0
	if s.valid() == true {
		t.Error("A special defense stat of 0 should not be valid.")
	}
}
