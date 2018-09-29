package game

// BattleGround play map stores max values of X,Y in axis
type BattleGround struct {
	MapSettings Coordinate
}

// Ship structure
type Ship struct {
	// Size of ship, how many cells it takes
	Size uint8
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
func NewBattleGround(x, y uint8) *BattleGround {
	return &BattleGround{
		MapSettings: Coordinate{
			AxisX: x,
			AxisY: y,
		},
	}
}
