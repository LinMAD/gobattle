package game

import (
	"testing"
)

func TestNewBattleGround(t *testing.T) {
	bg := NewBattleGround(5,5)
	if bg.MapSettings.AxisX != 5 || bg.MapSettings.AxisY != 5 {
		t.Error("New battle ground size not same as given")
	}
}

func TestBattleField_ValidateFleetCollision_Simple(t *testing.T) {
	getFleet := func(s1Loc, s2Loc Coordinate) []Ship {
		var s1, s2 Ship
		var fleet []Ship

		s1 = Ship{IsAlive:  true}
		s1.Location = append(s1.Location, s1Loc)
		s2 = Ship{IsAlive:  true}
		s2.Location = append(s2.Location, s2Loc)

		fleet = append(fleet, s1)
		fleet = append(fleet, s2)

		return fleet
	}
	bf := &BattleField{}

	// Validate collision on same place
	allZero := Coordinate{AxisX: 0, AxisY: 0}
	onSameErr := bf.ValidateFleetCollision(getFleet(allZero, allZero))
	if onSameErr == nil {
		t.Error("2 ship in fleet on same location, expected error")
	}

	// Validate collision check on top
	for i := 1; i <= 4; i++ {
		c1 := Coordinate{AxisX: uint8(i), AxisY: uint8(i)}
		c2 := Coordinate{AxisX: uint8(i+1), AxisY: uint8(i+1)}

		badFleet := bf.ValidateFleetCollision(getFleet(c1, c2))
		if badFleet == nil {
			t.Error("Expected error, because ships collides each other")
		}
	}
	// Validate collision check on bottom
	for i := 1; i <= 4; i++ {
		c1 := Coordinate{AxisX: uint8(i), AxisY: uint8(i)}
		c2 := Coordinate{AxisX: uint8(i-1), AxisY: uint8(i-1)}

		badFleet := bf.ValidateFleetCollision(getFleet(c1, c2))
		if badFleet == nil {
			t.Error("Expected error, because ships collides each other")
		}
	}
}

func TestBattleField_ValidateFleetCollision_Cornered(t *testing.T) {
	var isCollides error
	bf := &BattleField{}

	createShip := func(coordinates ...Coordinate) Ship {
		list := make([]Coordinate, 0)
		for _, v := range coordinates {
			list = append(list, v)
		}

		return Ship{IsAlive:  true, Location: list}
	}
	createFleet := func(ship ...Ship) []Ship {
		fleet := make([]Ship, 0)
		for _, s := range ship {
			fleet = append(fleet, s)
		}

		return fleet
	}

	s1 := createShip(Coordinate{AxisX: 5, AxisY: 5}, Coordinate{AxisX: 6, AxisY: 6})
	s2 := createShip(Coordinate{AxisX: 4, AxisY: 4})

	isCollides = bf.ValidateFleetCollision(createFleet(s1, s2))
	if isCollides == nil {
		t.Error("Expected error, 2 ships collides on corner on X4 and X5")
	}

	s3 := createShip(Coordinate{AxisX: 0, AxisY: 0}, Coordinate{AxisX: 1, AxisY: 0})
	s4 := createShip(Coordinate{AxisX: 2, AxisY: 1}, Coordinate{AxisX: 3, AxisY: 1})
	isCollides = bf.ValidateFleetCollision(createFleet(s3, s4))
	if isCollides == nil {
		t.Error("Expected error, 2 ships collides on corner on X1 and X2")
	}

	s5 := createShip(Coordinate{AxisX: 5, AxisY: 5}, Coordinate{AxisX: 5, AxisY: 6})
	s6 := createShip(Coordinate{AxisX: 3, AxisY: 3}, Coordinate{AxisX: 4, AxisY: 3})
	s7 := createShip(Coordinate{AxisX: 4, AxisY: 5})
	isCollides = bf.ValidateFleetCollision(createFleet(s5, s6, s7))
	if isCollides == nil {
		t.Error("Expected error, ships collides on corner on X3-4-5")
	}

	s8 := createShip(Coordinate{AxisX: 2, AxisY: 3})
	isCollides = bf.ValidateFleetCollision(createFleet(s2, s8))
	if isCollides == nil {
		t.Error("Expected error, 2 ships collides on corner on Y3 and Y4")
	}
}