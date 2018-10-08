package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"testing"
)

func TestStalkerStrategy_GetTargetLocation(t *testing.T) {
	sea := generator.NewSeaField(nil)

	// Put damaged ship
	sea[6][5] = game.GunMis
	sea[6][6] = game.GunMis
	sea[5][5] = game.GunHit
	sea[5][6] = game.GunMis
	sea[4][5] = game.GunMis
	sea[4][6] = game.GunMis

	stalker := stalkerStrategy{}
	target := stalker.GetTargetLocation(sea)
	if target.AxisY != 5 && target.AxisX != 4 {
		t.Error("Expected shot to last known location Y5, X4")
	}
}

func TestGridStrategy_GetTargetLocation(t *testing.T) {
	sea := generator.NewSeaField(nil)

	gd := &gridStrategy{}
	allMoves := 0
	for {
		target := gd.GetTargetLocation(sea)
		if target == nil {
			break
		}

		sea[target.AxisY][target.AxisX] = game.GunHit
		allMoves++
	}

	if allMoves != 24 {
		t.Error("Expected to be 24 moves in current strategy, but counted:", allMoves)
	}
}
