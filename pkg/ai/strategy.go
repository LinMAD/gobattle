package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"time"
)

type shotStrategy interface {
	GetTargetLocation(sea [][]string) *game.Coordinate
}

type (
	stalkerStrategy struct{}
	gridStrategy    struct{}
	randomStrategy  struct{}
)

// GetTargetLocation will try hit known ship
func (stalkerStrategy) GetTargetLocation(sea [][]string) *game.Coordinate {
	var y, x int8
	for y = 0; y < game.FSize; y++ {
		for x = 0; x < game.FSize; x++ {
			if sea[y][x] == game.GunHit {
				if x+1 < game.FSize && x-1 > 0 {
					if sea[y][x+1] == game.FNone {
						return &game.Coordinate{AxisX: x + 1, AxisY: y}
					}
					if sea[y][x-1] == game.FNone {
						return &game.Coordinate{AxisX: x - 1, AxisY: y}
					}
				}
				if y+1 < game.FSize && y-1 > 0 {
					if sea[y+1][x] == game.FNone {
						return &game.Coordinate{AxisX: x, AxisY: y + 1}
					}
					if sea[y-1][x] == game.FNone {
						return &game.Coordinate{AxisX: x, AxisY: y - 1}
					}
				}
			}
		}
	}

	return nil // Cant find scouted ship
}

// GetTargetLocation calculate totally random coordinate
func (randomStrategy) GetTargetLocation(sea [][]string) *game.Coordinate {
	st := time.Now()
	// TODO Make random more smarter, try hit location where less known area
	for {
		t := &game.Coordinate{
			AxisX: generator.RandomNum(0, int(game.FSize)-1),
			AxisY: generator.RandomNum(0, int(game.FSize)-1),
		}
		markOnField := sea[t.AxisY][t.AxisX]
		if markOnField != game.GunHit && markOnField != game.GunMis {
			return t
		}

		// try generate random target in second
		if time.Now().Sub(st) <= time.Second {
			continue
		}

		// ok, get any non shot location
		for y, xLine := range sea {
			for x := range xLine {
				if sea[y][x] != game.GunHit {
					return &game.Coordinate{AxisX: int8(x), AxisY: int8(y)}
				}
			}
		}
	}

	return nil
}

// GetTargetLocation scout whole sea in grid order
func (gridStrategy) GetTargetLocation(sea [][]string) *game.Coordinate {
	// help locate next unknown location in line
	isToNext := func(nextY int8, sea [][]string, target *game.Coordinate) bool {
		markOnField := sea[target.AxisY][target.AxisX]
		if markOnField == game.GunHit || markOnField == game.GunMis {
			target.AxisY = nextY
			markOnField = sea[target.AxisY][target.AxisX]
			if markOnField == game.GunHit || markOnField == game.GunMis {
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
