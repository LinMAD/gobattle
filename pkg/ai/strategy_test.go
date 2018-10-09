package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"github.com/LinMAD/gobattle/pkg/render"
	"testing"
)

func TestStalkerStrategy_GetTargetLocation(t *testing.T) {
	var target *game.Coordinate
	stalker := stalkerStrategy{}
	sea := generator.NewSeaField(nil)

	// Put vertical damaged ship
	sea[0][0] = game.GunHit
	sea[1][0] = game.GunHit
	target = stalker.GetTargetLocation(sea)
	if target.AxisY != 2 && target.AxisX != 0 {
		render.ShowBattleField(render.Screen{Title: "test", BattleField: sea}, true)
		t.Error("Unexpected target for shot, given", target)
	}
	sea[2][0] = game.GunMis

	// Put horizontal damaged ship
	sea[5][4] = game.GunHit
	sea[5][5] = game.GunHit
	sea[5][6] = game.GunMis
	target = stalker.GetTargetLocation(sea)
	if target.AxisY != 5 && target.AxisX != 3 {
		render.ShowBattleField(render.Screen{Title: "test", BattleField: sea}, true)
		t.Error("Unexpected target for shot, given", target)
	}
	sea[5][3] = game.GunHit
	sea[5][2] = game.GunMis

	sea[9][9] = game.GunHit
	sea[9][8] = game.GunHit
	target = stalker.GetTargetLocation(sea)
	if target.AxisY != 9 && target.AxisX != 7 {
		render.ShowBattleField(render.Screen{Title: "test", BattleField: sea}, true)
		t.Error("Unexpected target for shot, given", target)
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
