package game

import "fmt"

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
	// Location list of coordinates where ship located
	Location []Coordinate
}

// Coordinate are concrete position on X,Y axis
type Coordinate struct {
	// AxisX value of X in axis
	AxisX uint8
	// AxisY value of Y in axis
	AxisY uint8
}

// newBattleGround
func newBattleGround(x, y uint8) *BattleField {
	return &BattleField{
		MapSettings: Coordinate{
			AxisX: x,
			AxisY: y,
		},
	}
}

// ValidateFleetCollision return error with location if they collides
func (bf *BattleField) ValidateFleetCollision(fleet []Ship) error {
	if len(fleet) == 1 {
		return nil
	}

	for _, ship := range fleet {
		for _, nextShip := range fleet {
			if fmt.Sprint(&ship) == fmt.Sprint(&nextShip) {
				continue // Skip same ship comparision
			}

			if isCollides(&ship, &nextShip) {
				return fmt.Errorf(
					"collision found in fleet. Ship: %v and Ship: %v", ship.Location, nextShip.Location,
				)
			}
		}
	}

	return nil
}

// isCollides checks if ships encounter each other on top\bottom\same or even with corners
func isCollides(a *Ship, b *Ship) (isCollides bool) {
	for _, aShip := range a.Location {
		for _, bShip := range b.Location {
			// Check collision on right, left
			if int8(aShip.AxisX+1) == int8(bShip.AxisX) || int8(aShip.AxisX-1) == int8(bShip.AxisX) {
				return true
			}
			// Check collision on top, bottom
			if int8(aShip.AxisY+1) == int8(bShip.AxisY) || int8(aShip.AxisY-1) == int8(bShip.AxisY) {
				return true
			}
		}
	}

	return false
}
