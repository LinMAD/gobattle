package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"math/rand"
	"time"
)

type shotStrategy interface {
	GetTargetLocation(sea [][]int8) *game.Coordinate
}

type (
	gridStrategy   struct{}
	randomStrategy struct{}
)

// GetTargetLocation calculate totally random coordinate
func (randomStrategy) GetTargetLocation(sea [][]int8) *game.Coordinate {
	st := time.Now()
	rand.Seed(st.Unix())

	// TODO Make random more smarter, try hit location where less known area

	for {
		t := &game.Coordinate{
			AxisX: int8(rand.Intn(int(game.FSize) - 1)),
			AxisY: int8(rand.Intn(int(game.FSize) - 1)),
		}

		if sea[t.AxisY][t.AxisX] != game.FShot {
			return t
		}

		// if can't generate random target in second, get any non shot location
		if time.Now().Sub(st) >= time.Second {
			for y, xLine := range sea {
				for x := range xLine {
					if sea[y][x] != game.FShot {
						return &game.Coordinate{
							AxisX: int8(x),
							AxisY: int8(y),
						}
					}
				}
			}
		}
	}

	return nil
}

// GetTargetLocation scout whole sea in grid order
func (gridStrategy) GetTargetLocation(sea [][]int8) *game.Coordinate {
	// help locate next unknown location in line
	isToNext := func(nextY int8, sea [][]int8, target *game.Coordinate) bool {
		if sea[target.AxisY][target.AxisX] == game.FShot {
			target.AxisY = nextY
			if sea[target.AxisY][target.AxisX] == game.FShot {
				return true
			}
		}

		return false
	}

	// calculate target in Y line
	var y int8
	for y = 0; y <= game.FSize; y++ {
		switch y {
		case 0, 4, 8:
			target := &game.Coordinate{AxisX: y, AxisY: 3}
			if isToNext(7, sea, target) {
				continue
			}

			return target
		case 1, 5, 9:
			target := &game.Coordinate{AxisX: y, AxisY: 2}
			if isToNext(6, sea, target) {
				continue
			}

			return target
		case 2, 6:
			target := &game.Coordinate{AxisX: y}
			if isToNext(4, sea, target) {
				if isToNext(8, sea, target) {
					continue
				}
			}

			return target
		case 3, 7:
			target := &game.Coordinate{AxisX: y, AxisY: 1}
			if isToNext(5, sea, target) {
				if isToNext(9, sea, target) {
					continue
				}
			}

			return target
		}
	}

	return nil // All possible moves ended
}
