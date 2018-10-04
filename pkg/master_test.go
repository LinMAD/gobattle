package pkg

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"testing"
)

func TestNewGameAndShootShip(t *testing.T) {
	var gm *GameMaster
	var err error

	fleet := make([]*game.Ship, 0)

	_, err = NewGame("p1", fleet)
	if err == nil {
		t.Error("Expected error, fleet are empty")
	}

	shipCoordinate := game.Coordinate{AxisX: 0, AxisY: 0}
	shipLocation := make([]game.Coordinate, 1)
	shipLocation[0] = shipCoordinate
	fleet = append(fleet, &game.Ship{IsAlive: true, Location: shipLocation})
	gm, err = NewGame("p1", fleet)
	if err != nil {
		t.Error(err.Error())
	}

	if gm.room.GetActivePlayer().GetName() != "p1" {
		t.Error("Expected same player in room -> p1")
	}

	if gm.StillPlaying == false {
		t.Error("Expected game state still playing")
	}

	if gm.ShootInCoordinate(game.Coordinate{AxisX: 0, AxisY: 0}) == false {
		t.Error("Expected tru on hit of the ship")
	}
}