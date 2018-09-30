package game

import (
	"testing"
)

func TestNewBattleGround(t *testing.T) {
	bg := NewBattleGround(5, 5)
	if bg.MapSettings.AxisX != 5 || bg.MapSettings.AxisY != 5 {
		t.Error("New battle ground size not same as given")
	}
}

func TestBattleField_ValidateFleetCollision_Simple(t *testing.T) {
	getFleet := func(s1Loc, s2Loc Coordinate) []Ship {
		var s1, s2 Ship
		var fleet []Ship

		s1 = Ship{IsAlive: true}
		s1.Location = append(s1.Location, s1Loc)
		s2 = Ship{IsAlive: true}
		s2.Location = append(s2.Location, s2Loc)

		fleet = append(fleet, s1)
		fleet = append(fleet, s2)

		return fleet
	}
	bf := &BattleField{}

	// Validate collision check on top
	c1 := Coordinate{AxisX: 4, AxisY: 4}
	c2 := Coordinate{AxisX: 4, AxisY: 5}

	badFleet := bf.ValidateFleetCollision(getFleet(c1, c2))
	if badFleet == nil {
		t.Error("Expected error, because ships collides each other")
	}

	// Validate collision check on bottom
	c1 = Coordinate{AxisX: 4, AxisY: 4}
	c2 = Coordinate{AxisX: 4, AxisY: 3}

	badFleet = bf.ValidateFleetCollision(getFleet(c1, c2))
	if badFleet == nil {
		t.Error("Expected error, because ships collides each other")
	}
}

func TestBattleField_ValidateFleetCollision_Cornered(t *testing.T) {
	var isCollides error
	bf := &BattleField{}

	s1 := helpCreateShip(Coordinate{AxisX: 5, AxisY: 5})
	s2 := helpCreateShip(Coordinate{AxisX: 4, AxisY: 4})

	isCollides = bf.ValidateFleetCollision(helpCreateFleet(s1, s2))
	if isCollides == nil {
		t.Error("Expected error, 2 ships collides on corner")
	}

	s3 := helpCreateShip(Coordinate{AxisX: 5, AxisY: 7})
	s4 := helpCreateShip(Coordinate{AxisX: 6, AxisY: 6})
	isCollides = bf.ValidateFleetCollision(helpCreateFleet(s3, s4))
	if isCollides == nil {
		t.Error("Expected error, 2 ships collides on corner")
	}

	s5 := helpCreateShip(Coordinate{AxisX: 3, AxisY: 3}, Coordinate{AxisX: 3, AxisY: 4})
	s6 := helpCreateShip(Coordinate{AxisX: 4, AxisY: 2}, Coordinate{AxisX: 4, AxisY: 1})
	s7 := helpCreateShip(Coordinate{AxisX: 2, AxisY: 2})
	isCollides = bf.ValidateFleetCollision(helpCreateFleet(s5, s6, s7))
	if isCollides == nil {
		t.Error("Expected error, ships collides on corner on X3-4-5")
	}

	s8 := helpCreateShip(Coordinate{AxisX: 2, AxisY: 3})
	isCollides = bf.ValidateFleetCollision(helpCreateFleet(s2, s8))
	if isCollides == nil {
		t.Error("Expected error, 2 ships collides on corner on Y3 and Y4")
	}
}

func TestBattleField_ValidateFleetCollision_CorrectFleet(t *testing.T) {
	l1 := Coordinate{0, 0}
	l2 := Coordinate{2, 2}
	l3 := Coordinate{0, 4}
	l4 := Coordinate{6, 2}

	fleet := helpCreateFleet(
		helpCreateShip(l1),
		helpCreateShip(l2),
		helpCreateShip(l3),
		helpCreateShip(l4),
	)

	bf := &BattleField{}
	isCollided := bf.ValidateFleetCollision(fleet)
	if isCollided != nil {
		t.Error("Expected to be correct fleet with no collision")
	}
}

func helpCreateShip(coordinates ...Coordinate) Ship {
	list := make([]Coordinate, 0)
	for _, v := range coordinates {
		list = append(list, v)
	}

	return Ship{IsAlive: true, Location: list}
}

func helpCreateFleet(ship ...Ship) []Ship {
	fleet := make([]Ship, 0)
	for _, s := range ship {
		fleet = append(fleet, s)
	}

	return fleet
}