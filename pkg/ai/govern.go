package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
)

// Govern AI\Bot
type Govern struct {
	name string
	// seaPlan whole picture of sea
	seaPlan [][]string
}

// NewGovern
func NewGovern() *Govern {
	return &Govern{
		name:    "Govern",
		seaPlan: generator.NewSeaField(nil),
	}
}

// GetName of bot
func (g *Govern) GetName() string {
	return g.name
}

// GetSeaPlan bot battle field with sea and enemy ships
func (g *Govern) GetSeaPlan() [][]string {
	return g.seaPlan
}

// OpenFire decide target where to shoot
func (g *Govern) OpenFire() game.Coordinate {
	stalker := stalkerStrategy{}
	grid := gridStrategy{}
	rand := randomStrategy{}

	target := stalker.GetTargetLocation(g.seaPlan)
	if target == nil {
		target = grid.GetTargetLocation(g.seaPlan)
	}
	if target == nil {
		target = rand.GetTargetLocation(g.seaPlan)
	}

	return *target
}

// CollectResultOfShot
func (g *Govern) CollectResultOfShot(t game.Coordinate, isHit bool) {
	if isHit == false {
		g.seaPlan[t.AxisY][t.AxisX] = game.GunMis

		return
	}

	g.seaPlan[t.AxisY][t.AxisX] = game.GunHit
}
