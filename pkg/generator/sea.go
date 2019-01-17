package generator

import "github.com/LinMAD/gobattle/pkg/game"

// NewSeaField create sea map
func NewSeaField(fleet []*game.Ship) [][]string {
	seaPlan := make([][]string, game.FSize)

	// Create sea plan
	for y := 0; y < game.FSize; y++ {
		line := make([]string, game.FSize)
		for x := 0; x < game.FSize; x++ {
			line[x] = game.FNone
		}
		seaPlan[y] = line
	}

	if fleet == nil {
		return seaPlan
	}

	for _, s := range fleet {
		for _, coordinate := range s.Location {
			seaPlan[coordinate.AxisY][coordinate.AxisX] = game.FShip
		}
	}

	return seaPlan
}
