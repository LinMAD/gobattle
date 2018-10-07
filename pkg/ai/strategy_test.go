package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"testing"
)

func TestGridStrategy_GetTargetLocation(t *testing.T) {
	sea := generator.GenerateSeaPlan(nil)

	gd := &gridStrategy{}
	allMoves := 0
	for {
		target := gd.GetTargetLocation(sea)
		if target == nil {
			break
		}

		sea[target.AxisY][target.AxisX] = game.FShot
		allMoves++
	}

	if allMoves != 24 {
		t.Error("Expected to be 24 moves in current strategy, but counted:", allMoves)
	}
}
