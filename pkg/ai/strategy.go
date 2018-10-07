package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
)

type shotStrategy interface {
	GetTargetLocation(sea [][]int8) *game.Coordinate
}

type (
	gridStrategy struct{}
)

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
			target := &game.Coordinate{AxisX:y, AxisY:3}
			if isToNext(7, sea, target) {
				continue
			}

			return target
		case 1, 5, 9:
			target := &game.Coordinate{AxisX:y, AxisY:2}
			if isToNext(6, sea, target) {
				continue
			}

			return target
		case 2, 6:
			target := &game.Coordinate{AxisX:y}
			if isToNext(4, sea, target) {
				if isToNext(8, sea, target) {
					continue
				}
			}

			return target
		case 3, 7:
			target := &game.Coordinate{AxisX:y, AxisY:1}
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
