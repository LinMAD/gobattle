package generator

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/game"
	"testing"
)

func TestFleetFromString_Empty(t *testing.T) {
	strFleet := ""

	_, err := FleetFromString(strFleet)
	if err == nil {
		t.Error("Expected error, fleet string are empty")
	}
}

func TestFleetFromString_OneCellShips(t *testing.T) {
	strFleet := "01-00-10"

	fleet, err := FleetFromString(strFleet)
	handleCommonError(err, t)

	c1 := game.Coordinate{AxisX: 1, AxisY: 0}
	c2 := game.Coordinate{AxisX: 0, AxisY: 2}

	err = validateLocationFleet(2, fleet, c1, c2)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFleetFromString_TwoCellShip(t *testing.T) {
	strFleet := "11-00-11"

	fleet, err := FleetFromString(strFleet)
	handleCommonError(err, t)

	c1 := game.Coordinate{AxisX: 0, AxisY: 0}
	c2 := game.Coordinate{AxisX: 1, AxisY: 0}
	c3 := game.Coordinate{AxisX: 0, AxisY: 2}
	c4 := game.Coordinate{AxisX: 1, AxisY: 2}

	err = validateLocationFleet(2, fleet, c1, c2, c3, c4)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFleetFromString_OneShip(t *testing.T) {
	strFleet := "1"

	fleet, err := FleetFromString(strFleet)
	handleCommonError(err, t)

	c1 := game.Coordinate{AxisX: 0, AxisY: 0}

	err = validateLocationFleet(1, fleet, c1)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFleetFromString_VerticalShipTwoCells(t *testing.T) {
	strFleet := "01-01"

	fleet, err := FleetFromString(strFleet)
	handleCommonError(err, t)

	c1 := game.Coordinate{AxisX: 1, AxisY: 0}
	c2 := game.Coordinate{AxisX: 1, AxisY: 1}

	err = validateLocationFleet(1, fleet, c1, c2)
	if err == nil {
		t.Errorf("Expected to be error, currently vertical ship parsing not supported")
	}
}

func handleCommonError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected to be converted string to slice, err: %v", err.Error())
	}
}

func validateLocationFleet(expectedShips int, fleet []game.Ship, expectedCoordinates ...game.Coordinate) error {
	if len(fleet) != expectedShips {
		return fmt.Errorf("expected size of the fleet %d (ships) but it's: %d (ships)", expectedShips, len(fleet))
	}

	isFoundExpectedLocations := make([]bool, len(expectedCoordinates))
	for _, ship := range fleet {
		for _, loc := range ship.Location {
			for i, expected := range expectedCoordinates {
				if expected == loc {
					isFoundExpectedLocations[i] = true
				}
			}
		}
	}

	for _, isFound := range isFoundExpectedLocations {
		if !isFound {
			return fmt.Errorf("expected to be found all locations in ships")
		}
	}

	return nil
}
