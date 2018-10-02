package game

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	room := NewWarRoom()

	fleet := make([]*Ship, 0)
	name := "Gopher"
	p, err := NewPlayer(name, fleet, room)

	if err == nil {
		t.Error("Expected to be error, player fleet empty")
	}

	fleet = append(fleet, &Ship{})
	p, err = NewPlayer(name, fleet, room)
	if err == nil {
		t.Error("Expected error, ship missed location")
	}

	fleet = make([]*Ship, 0)
	fleet = append(fleet, &Ship{})
	fleet[0].Location = append(fleet[0].Location, Coordinate{AxisX: 5, AxisY: 5})
	p, err = NewPlayer(name, fleet, room)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}

	if p.GetName() != name {
		t.Error("Player created incorrect, name not same -> " + name)
	}

	if len(p.GetFleet()) == 0 {
		t.Error("Expected empty list of ships for new player")
	}
}

func TestNewPlayer_ValidateFleetCollision(t *testing.T) {
	room := NewWarRoom()

	fleet := helpCreateFleet(
		helpCreateShip(Coordinate{1, 1}),
		helpCreateShip(Coordinate{2, 2}),
	)

	_, err := NewPlayer("P1", fleet, room)
	if err == nil {
		t.Error("Expected to be error, incorrect fleet")
	}
	if err.Error() != "collision found in fleet. Ship: [{1 1}] and Ship: [{2 2}]" {
		t.Error("Unexpected error -> ", err.Error())
	}
}
