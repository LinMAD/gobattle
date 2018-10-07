package ai

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
)

// Govern AI\Bot
type Govern struct {
	name string
	// seaPlan whole picture of sea
	seaPlan [][]int8
	// foeFleet stores located fleet
	foeFleet []*game.Ship
}

// NewGovern
func NewGovern() *Govern {
	return &Govern{
		name:    "Govern",
		seaPlan: generator.GenerateSeaPlan(nil),
	}
}

// GetName of bot
func (g *Govern) GetName() string {
	return g.name
}

// OpenFire decide target where to shoot
func (g *Govern) OpenFire() game.Coordinate {
	grid := gridStrategy{}
	// TODO Add strategy with remembered ships
	// TODO Get random strategy between random and grid
	t := grid.GetTargetLocation(g.seaPlan)
	if t == nil {
		rand := randomStrategy{}
		t = rand.GetTargetLocation(g.seaPlan)
	}

	g.seaPlan[t.AxisY][t.AxisX] = game.FShot

	return *t
}
