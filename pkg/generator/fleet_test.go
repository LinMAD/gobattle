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

	ship = createShipRandom(1)
	if len(ship.Location) != 1 {
		t.Error("Ships must be size of ", 1)
	}
	if createShipRandom(1).Location[0] == ship.Location[0] {
		t.Error("New ship has same location as before")
	}

	ship = createShipRandom(2)
	if len(ship.Location) != 2 {
		t.Error("Ships must be size of ", 2)
	}
	for i, l := range createShipRandom(2).Location {
		if l == ship.Location[i] {
			t.Error("Location same as before")
		}
	}
}

func BenchmarkNewFleet(b *testing.B) {
	NewFleet()
}

func BenchmarkNew10FleetsInRoutine(b *testing.B) {
	for i := 0; i < 10; i++ {
		NewFleet()
	}
}

func benchmarkCreateShip(size uint8, b *testing.B) {
	for n := 0; n < b.N; n++ {
		createShipRandom(size)
	}
}

func BenchmarkCreateShipSize1(b *testing.B) {
	benchmarkCreateShip(1, b)
}

func BenchmarkCreateShipSize2(b *testing.B) {
	benchmarkCreateShip(2, b)
}

func BenchmarkCreateShipSize3(b *testing.B) {
	benchmarkCreateShip(3, b)
}

func BenchmarkCreateShipSize4(b *testing.B) {
	benchmarkCreateShip(4, b)
}

func BenchmarkCreateShipSize5(b *testing.B) {
	benchmarkCreateShip(5, b)
}

func BenchmarkCreateShipSize6(b *testing.B) {
	benchmarkCreateShip(6, b)
}

