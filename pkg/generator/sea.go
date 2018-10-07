package generator

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"sync"
)

// NewSeaField create sea map
func NewSeaField(fleet []*game.Ship) [][]string {
	seaPlan := make([][]string, game.FSize)

	// Create sea plan
	for y := 0; int8(y) < game.FSize; y++ {
		line := make([]string, game.FSize)
		for x := 0; int8(x) < game.FSize; x++ {
			line[x] = game.FNone
		}
		seaPlan[y] = line
	}

	if fleet == nil {
		return seaPlan
	}

	var w int
	var wg sync.WaitGroup
	for w = 0; w < len(fleet); w++ {
		wg.Add(1)
		go func(w int, seaPlan [][]string) {
			defer wg.Done()
			for _, coordinate := range fleet[w].Location {
				seaPlan[coordinate.AxisY][coordinate.AxisX] = game.FShip
			}
		}(w, seaPlan)
	}

	wg.Wait()

	return seaPlan
}
