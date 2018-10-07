package pkg

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"testing"
)

func TestNewGameAndShootShip(t *testing.T) {
	var gm *GameMaster
	var err error
	var settings GameSettings

	fleet := make([]*game.Ship, 0)

	settings.PlayerName = "p1"
	settings.PlayerFleet = fleet
	_, err = NewGame(settings)
	if err == nil {
		t.Error("Expected error, fleet are empty")
	}

	shipCoordinate := game.Coordinate{AxisX: 0, AxisY: 0}
	shipLocation := make([]game.Coordinate, 1)
	shipLocation[0] = shipCoordinate
	fleet = append(fleet, &game.Ship{IsAlive: true, Location: shipLocation})
	settings.PlayerFleet = fleet
	gm, err = NewGame(settings)
	if err != nil {
		t.Error(err.Error())
	}

	if gm.room.GetActivePlayer().GetName() != "p1" {
		t.Error("Expected same player in room -> p1")
	}

	if gm.StillPlaying == false {
		t.Error("Expected game state still playing")
	}
}
