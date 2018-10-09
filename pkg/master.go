package pkg

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/ai"
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"github.com/LinMAD/gobattle/pkg/render"
	"time"
)

// GameSettings for new game
type GameSettings struct {
	// PlayerName on display
	PlayerName string
	// PlayerFleet collection of ships
	PlayerFleet []*game.Ship
	// IsVersusHuman flag to show screens
	IsVersusHuman bool
	// GameSpeed how fast changes moves (for bots) in millisecond
	GameSpeed time.Duration
}

// GameMaster acts mostly like proxy to follow the game rules, encapsulates mechanics
type GameMaster struct {
	// versusHuman
	versusHuman bool
	// room with players
	room *game.WarRoom
	// StillPlaying until one of the fleet not destroyed
	StillPlaying bool
	// GameEndReason if game ends then will be known reason of ending
	GameEndReason string
	// bot to emulate gaming
	bot *ai.Govern
	// timeToSleep to hold game speed in millisecond
	timeToSleep time.Duration
}

// NewGame creates game to play
func NewGame(settings GameSettings) (*GameMaster, error) {
	var errPlayer error
	gp := &GameMaster{
		versusHuman:  settings.IsVersusHuman,
		room:         game.NewWarRoom(),
		StillPlaying: true,
		bot:          ai.NewGovern("Govern"),
		timeToSleep:  settings.GameSpeed,
	}

	_, errPlayer = game.NewPlayer(settings.PlayerName, settings.PlayerFleet, gp.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	_, errPlayer = game.NewPlayer(gp.bot.GetName(), generator.NewFleet(), gp.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	return gp, nil
}

// ShootInCoordinate
func (gp *GameMaster) ShootInCoordinate(target game.Coordinate) bool {
	var isFirstPlayerShot bool

	// Here can be tricky
	// if next turn not a build in bot then
	// return only if shot hit a ship
	nextPlayer := gp.room.GetActivePlayer()
	if nextPlayer.GetName() != gp.bot.GetName() {
		isFirstPlayerShot = nextPlayer.GunShoot(target)
		gp.checkPlayerFleet(gp.room.GetOppositePlayer(nextPlayer.GetName()))
		if isFirstPlayerShot {
			return isFirstPlayerShot
		}
	}

	// So first player loosed initiative
	// Now bot will try shot all ships while first player waits
	nextPlayer = gp.room.GetActivePlayer()
	oppositePlayer := gp.room.GetOppositePlayer(nextPlayer.GetName())
	for {
		// AI did action
		targetToHit := gp.bot.OpenFire()
		isBotHit := nextPlayer.GunShoot(targetToHit)
		// Collect result for bot
		gp.bot.CollectResultOfShot(targetToHit, isBotHit)
		// Check if bot win the game
		gp.checkPlayerFleet(oppositePlayer)
		if gp.versusHuman == false {
			render.ShowBattleField(
				render.Screen{
					Title:       "Shooting field of " + gp.bot.GetName(),
					BattleField: gp.bot.GetSeaPlan(),
				},
				true,
			)
			time.Sleep(gp.timeToSleep * time.Millisecond) // Slow down game speed
		}

		// Ok if bot win or miss then return back control to player
		if gp.StillPlaying == false || isBotHit == false {
			break
		}
	}

	return isFirstPlayerShot
}

// checkPlayerFleet check if game ended
func (gp *GameMaster) checkPlayerFleet(p *game.Player) {
	fleetSize := len(p.GetFleet())

	for _, s := range p.GetFleet() {
		if s.IsAlive == false {
			fleetSize--
		}
	}

	if fleetSize == 0 {
		gp.StillPlaying = false
		gp.GameEndReason = fmt.Sprintf("%s lossed whole fleet", p.GetName())

		return
	}
}
