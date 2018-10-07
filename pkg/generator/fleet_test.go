package generator

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"testing"
)

func TestNewFleet(t *testing.T) {
	fleet := NewFleet()
	err := game.ValidateFleetCollision(fleet)
	if err != nil {
		t.Error(err.Error())
	}
	if len(fleet) != len(game.ShipTypes) {
		t.Error("Expected same size of fleet as count of ship types")
	}
}

func TestCreateShip(t *testing.T) {
	var ship *game.Ship

	ship = createShip(1)
	if len(ship.Location) != 1 {
		t.Error("Ships must be size of ", 1)
	}
	if createShip(1).Location[0] == ship.Location[0] {
		t.Error("New ship has same location as before")
	}

	ship = createShip(2)
	if len(ship.Location) != 2 {
		t.Error("Ships must be size of ", 2)
	}
	for i, l := range createShip(2).Location {
		if l == ship.Location[i] {
			t.Error("Location same as before")
		}
	}
}
