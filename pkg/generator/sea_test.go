package generator

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"testing"
)

func TestGenerateSeaPlan(t *testing.T) {
	shipLocation := make([]game.Coordinate, 6)
	shipLocation[0] = game.Coordinate{AxisX: 5, AxisY: 2}
	shipLocation[1] = game.Coordinate{AxisX: 5, AxisY: 3}
	shipLocation[2] = game.Coordinate{AxisX: 5, AxisY: 4}
	shipLocation[3] = game.Coordinate{AxisX: 2, AxisY: 9}
	shipLocation[4] = game.Coordinate{AxisX: 1, AxisY: 9}
	shipLocation[5] = game.Coordinate{AxisX: 0, AxisY: 9}
	fleet := make([]*game.Ship, 0)
	fleet = append(fleet, &game.Ship{Location: shipLocation})

	seaPlan := GenerateSeaPlan(fleet)
	if seaPlan == nil {
		t.Error("sea plan expected to be returned")
	}

	for _, expected := range shipLocation {
		_ = seaPlan[expected.AxisY][expected.AxisY]
	}
}
