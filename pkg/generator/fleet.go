package generator

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/game"
)

// delimiterRune in string with given fleet like: 01-01 -> "-" represents part
const (
	delimiterRune = 45 // "-"
	shipRune      = 49 // "1"
)

// FleetFromString converts string to list of ships with X,Y coordinates
func FleetFromString(strFleet string) (fleet []game.Ship, err error) {
	if len(strFleet) == 0 {
		return nil, fmt.Errorf("empty string of fleet given")
	}

	// Parse string and create fleet
	var yCoordinate, xCoordinate uint8
	fleet = make([]game.Ship, 0)
	locations := make([]game.Coordinate, 0)

	// TODO Add vertical ship parsing, like parse string to matrix then search locations
	for strIndex, chrRune := range strFleet {
		if chrRune == delimiterRune {
			xCoordinate = 0
			locations = make([]game.Coordinate, 0)

			yCoordinate++

			continue
		}

		// Locate ships and build fleet
		if chrRune != shipRune {
			xCoordinate++

			continue // Current field not a ship
		}

		// If next is a ship, then add current location and continue parsing
		if len(strFleet) > strIndex+1 {
			if strFleet[strIndex+1] == shipRune {
				c := game.Coordinate{AxisX: uint8(xCoordinate), AxisY: yCoordinate}
				locations = append(locations, c)
				xCoordinate++

				continue
			}
		}

		// Add ship to fleet with coordinates
		c := game.Coordinate{AxisX: uint8(xCoordinate), AxisY: yCoordinate}
		locations = append(locations, c)

		s := game.Ship{IsAlive: true, Location: locations}
		fleet = append(fleet, s)
		xCoordinate++
	}

	return fleet, nil
}
