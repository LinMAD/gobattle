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
	// foeFleet known ships in sea
	foeFleet []*game.Ship
}

// NewGovern
func NewGovern() *Govern {
	bot := &Govern{
		name:     "Govern",
		seaPlan:  generator.NewSeaField(nil),
		foeFleet: make([]*game.Ship, 0),
	}

	for s := 0; s < len(game.ShipTypes); s++ {
		bot.foeFleet = append(bot.foeFleet, &game.Ship{IsAlive: true})
	}

	return bot
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
	grid := gridStrategy{}
	// TODO Add strategy with remembered ships
	// TODO Get random strategy between random and grid
	t := grid.GetTargetLocation(g.seaPlan)
	if t == nil {
		rand := randomStrategy{}
		t = rand.GetTargetLocation(g.seaPlan)
	}

	return *t
}

// CollectResultOfShot
func (g *Govern) CollectResultOfShot(t game.Coordinate, isHit bool) {
	if isHit == false {
		g.seaPlan[t.AxisY][t.AxisX] = game.GunMis

		return
	}

	g.seaPlan[t.AxisY][t.AxisX] = game.GunHit
	for _, scoutedShip := range g.foeFleet {
		if len(scoutedShip.DamagedLocation) == 0 {
			scoutedShip.DamagedLocation = append(scoutedShip.DamagedLocation, &t)

			return
		}

		for _, loc := range scoutedShip.DamagedLocation {
			if loc.AxisY == t.AxisY {
				if loc.AxisX+1 == t.AxisX || loc.AxisX-1 == t.AxisX {
					scoutedShip.DamagedLocation = append(scoutedShip.DamagedLocation, &t)
					return
				}
			}
			if loc.AxisX == t.AxisX {
				if loc.AxisY+1 == t.AxisY || loc.AxisY-1 == t.AxisY {
					scoutedShip.DamagedLocation = append(scoutedShip.DamagedLocation, &t)
				}
			}
		}
	}
}
