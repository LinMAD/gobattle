package game

import "fmt"

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

// NewBattleGround
func NewBattleGround(x, y uint8) *BattleField {
	return &BattleField{
		MapSettings: Coordinate{
			AxisX: x,
			AxisY: y,
		},
	}
}

// ValidateFleetCollision return error with location if they collides
func (bf *BattleField) ValidateFleetCollision(fleet []Ship) error {
	if len(fleet) <= 1 {
		return nil
	}

	for _, ship := range fleet {
		for _, nextShip := range fleet {
			if isCollides(ship, nextShip) {
				return fmt.Errorf(
					"collision found in fleet. Ship: %v and Ship: %v", ship.Location, nextShip.Location,
				)
			}
		}
	}

	return nil
}

// isCollides checks if ships encounter each other on top\bottom\same or even with corners
func isCollides(a Ship, b Ship) (isCollides bool) {
	for _, aShip := range a.Location {
		for _, bShip := range b.Location {
			if aShip.AxisX == bShip.AxisX && aShip.AxisY != bShip.AxisY {
				return true
			}
			if aShip.AxisX != bShip.AxisX && aShip.AxisY == bShip.AxisY {
				return true
			}
			if aShip.AxisX == bShip.AxisX && aShip.AxisY == bShip.AxisY {
				return true
			}

			if int8(aShip.AxisX + 1) == int8(bShip.AxisX) || int8(aShip.AxisX + 2) == int8(bShip.AxisX)  {
				return true
			}
			if int8(aShip.AxisX - 1) == int8(bShip.AxisX) || int8(aShip.AxisX - 2) == int8(bShip.AxisX)  {
				return true
			}

			if int8(aShip.AxisY + 1) == int8(bShip.AxisY) || int8(aShip.AxisY + 2) == int8(bShip.AxisY)  {
				return true
			}
			if int8(aShip.AxisY - 1) == int8(bShip.AxisY) || int8(aShip.AxisY - 2) == int8(bShip.AxisY)  {
				return true
			}
		}
	}

	return false
}