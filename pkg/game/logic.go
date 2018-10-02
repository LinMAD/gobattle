package game

import (
	"fmt"
)

// TODO Think about split game logic to sub package: field, ships

const (
	// General runes in communication, ships locations and part separator
	MsgDelimiter rune = 45 // "-"
	MsgNone      rune = 48 // "0"
	MsgShip      rune = 49 // "1"
)

// BattleGround play map stores max values of X,Y in axis
type BattleField struct {
	MapSettings Coordinate
}

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

// newBattleGround
func newBattleGround(x, y int8) *BattleField {
	return &BattleField{MapSettings: Coordinate{AxisX: x, AxisY: y}}
}

// ValidateFleetCollision return error with location if they collides
func (bf *BattleField) ValidateFleetCollision(fleet []*Ship) error {
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
			s.Location[i] = Coordinate{-1, -1}
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
