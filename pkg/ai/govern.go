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

func (g *Govern) GetName() string {
	return g.name
}

func (g *Govern) OpenFire() game.Coordinate {
	// TODO Add more strategies to bombard player
	s := gridStrategy{}
	t := s.GetTargetLocation(g.seaPlan)
	if t == nil {
		// TODO Go random always
		t = &game.Coordinate{
			AxisX: 0,
			AxisY: 0,
		}
	}

	g.seaPlan[t.AxisY][t.AxisX] = game.FShot

	print("\n")
	for y := 9; y >= 0; y-- {
		for x := 0; x < 10; x++ {
			print(" ")
			print(g.seaPlan[y][x])
			print(" ")
		}
		print("\n")
	}
	print("\n")

	return *t
}

//
//func (g *Govern) CollectShotResult(isHit bool, target game.Coordinate) {
//	for _, s := range g.foeFleet {
//		for _, l := range s.Location {
//			if l == target
//		}
//	}
//}