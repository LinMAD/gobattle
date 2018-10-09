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

	// Collect all damaged fields
	for y = 0; y < game.FSize; y++ {
		for x = 0; x < game.FSize; x++ {
			if sea[y][x] != game.GunHit {
				continue
			}

			target := correctShotAccuracy(sea, &game.Coordinate{AxisX: int8(x), AxisY: int8(y)}, 1)
			if target != nil {
				return target
			}
		}
	}

	return nil // Cant find scouted ship
}

// GetTargetLocation calculate totally random coordinate
func (randomStrategy) GetTargetLocation(sea [][]string) *game.Coordinate {
	st := time.Now()

	for {
		// try generate random target in second
		if time.Now().Sub(st) > time.Second {
			// ok, get any location where can shoot
			for y := range sea {
				for x := range sea {
					if sea[y][x] != game.GunHit {
						return &game.Coordinate{AxisX: int8(y), AxisY: int8(y)}
					}
				}
			}
		}

		t := &game.Coordinate{
			AxisX: generator.RandomNum(0, int(game.FSize)-1),
			AxisY: generator.RandomNum(0, int(game.FSize)-1),
		}
		tc := correctShotAccuracy(sea, &game.Coordinate{AxisX: t.AxisX, AxisY: t.AxisY}, 1)
		if tc != nil {
			return tc
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

			return correctShotAccuracy(sea, target, 1)
		case 1, 5, 9:
			target := &game.Coordinate{AxisX: y, AxisY: 2}
			if isToNext(6, sea, target) {
				continue
			}

			return correctShotAccuracy(sea, target, 1)
		case 2, 6:
			target := &game.Coordinate{AxisX: y}
			if isToNext(4, sea, target) {
				if isToNext(8, sea, target) {
					continue
				}
			}

			return correctShotAccuracy(sea, target, 1)
		case 3, 7:
			target := &game.Coordinate{AxisX: y, AxisY: 1}
			if isToNext(5, sea, target) {
				if isToNext(9, sea, target) {
					continue
				}
			}

			return correctShotAccuracy(sea, target, 1)
		}
	}

	return nil // All possible moves ended
}

// correctShotAccuracy returns is target accurate shot
func correctShotAccuracy(sea [][]string, t *game.Coordinate, diff int8) *game.Coordinate {
	var nextMark, currentMark string

	// Locate how ship placed
	isShipVertical := false
	if t.AxisY+diff < game.FSize {
		currentMark = sea[t.AxisY][t.AxisX]
		nextMark = sea[t.AxisY+diff][t.AxisX]
		if currentMark == game.GunHit && nextMark == game.GunHit {
			isShipVertical = true
		}
	}
	if t.AxisY-diff >= 0 {
		currentMark = sea[t.AxisY][t.AxisX]
		nextMark = sea[t.AxisY-diff][t.AxisX]
		if currentMark == game.GunHit && nextMark == game.GunHit {
			isShipVertical = true
		}
	}

	// Calibrate by vertical\horizontal target
	switch isShipVertical {
	case true: // vertical target
		if t.AxisY+diff < game.FSize {
			if sea[t.AxisY+diff][t.AxisX] == game.FNone {
				t.AxisY += diff

				return t
			}
		}
		if t.AxisY-diff >= 0 {
			if sea[t.AxisY-diff][t.AxisX] == game.FNone {
				t.AxisY -= diff

				return t
			}
		}
	case false: // horizontal target
		if t.AxisX+diff < game.FSize {
			if sea[t.AxisY][t.AxisX+diff] == game.FNone {
				t.AxisX += diff

				return t
			}
		}
		if t.AxisX-diff >= 0 {
			if sea[t.AxisY][t.AxisX-diff] == game.FNone {
				t.AxisX -= diff

				return t
			}
		}
	}

	return nil
}
