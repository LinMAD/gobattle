package game

import (
	"fmt"
)

const (
	// Field marks
	FShot int8 = -1 // shot cell or ships
	FNone int8 = 0  // unknown field, empty
	FShip int8 = 1  // ship on field
	FSize int8 = 10 // size of battle field, max of X and Y axis (count from 0)
)

// ShipTypes in fleet, each represent one ship
var ShipTypes = []uint8{6, 5, 4, 3, 2, 1}

// Ship structure
type Ship struct {
	// IsAlive status whole ship
	IsAlive bool
	// Location list of coordinates where ship located, if it has negative coordinate, then ship is damaged
	Location []Coordinate
}

// Coordinate are concrete position on X,Y axis
type Coordinate struct {
	// AxisX value of X in axis
	AxisX int8
	// AxisY value of Y in axis
	AxisY int8
}

// ValidateFleetCollision return error with location if they collides
func ValidateFleetCollision(fleet []*Ship) error {
	if len(fleet) == 1 {
		return nil
	}

	for _, ship := range fleet {
		for _, nextShip := range fleet {
			if ship == nextShip {
				continue // Skip same ship comparision
			}

			if ship.isCollides(nextShip) {
				return fmt.Errorf(
					"collision found in fleet. Ship: %v and Ship: %v", ship.Location, nextShip.Location,
				)
			}
		}
	}

	return nil
}

// hitShip to kill or damage targeted ship by firing coordinates
func (s *Ship) isHit(firingCoordinate *Coordinate) bool {
	for i, shipLocation := range s.Location {
		if shipLocation.AxisX == firingCoordinate.AxisX && shipLocation.AxisY == firingCoordinate.AxisY {
			s.Location[i] = Coordinate{FShot, FShot}
			s.isStillAlive()

			return true
		}
	}

	return false
}

// isStillAlive check ship health by validating his location
func (s *Ship) isStillAlive() bool {
	var health, sizeOfShip int
	sizeOfShip = len(s.Location)
	health = sizeOfShip

	for _, l := range s.Location {
		if l.AxisY < 0 && l.AxisX < 0 {
			health--
		}
	}

	if health == 0 {
		s.IsAlive = false
	}

	return s.IsAlive
}

// isCollides checks if ships encounter each other on top\bottom\same or even with corners
func (s *Ship) isCollides(nextShip *Ship) (isCollides bool) {
	for _, aShip := range s.Location {
		for _, bShip := range nextShip.Location {
			// Check collision on right, left
			if aShip.AxisY == bShip.AxisY {
				if aShip.AxisX+1 == bShip.AxisX {
					return true
				}
				if aShip.AxisX-1 == bShip.AxisX {
					return true
				}
			}

			// Check collision on top, bottom
			if aShip.AxisX == bShip.AxisX {
				if aShip.AxisY+1 == bShip.AxisY {
					return true
				}
				if aShip.AxisY-1 == bShip.AxisY {
					return true
				}
			}

			// Check collision on corners on right
			if aShip.AxisX+1 == bShip.AxisX && aShip.AxisY+1 == bShip.AxisY {
				return true
			}
			if aShip.AxisX+1 == bShip.AxisX && aShip.AxisY-1 == bShip.AxisY {
				return true
			}
			// Check collision on corners on left
			if aShip.AxisX-1 == bShip.AxisX && aShip.AxisY+1 == bShip.AxisY {
				return true
			}
			if aShip.AxisX-1 == bShip.AxisX && aShip.AxisY-1 == bShip.AxisY {
				return true
			}
		}
	}

	return false
}
