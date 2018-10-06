package generator

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"sync"
)

func GenerateSeaPlan(fleet []*game.Ship) [][]int8 {
	seaPlan := make([][]int8, game.FSize)

	// Create sea plan
	for y := 0; int8(y) < game.FSize; y++ {
		line := make([]int8, game.FSize)
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
		go func(w int, seaPlan [][]int8) {
			defer wg.Done()
			for _, coordinate := range fleet[w].Location {
				seaPlan[coordinate.AxisY][coordinate.AxisX] = game.FShip
			}
		}(w, seaPlan)
	}

	wg.Wait()

	return seaPlan
}
