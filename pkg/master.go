package pkg

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/ai"
	"github.com/LinMAD/gobattle/pkg/game"
)

// GameMaster encapsulates game iterations and follows rules
type GameMaster struct {
	// room with players
	room *game.WarRoom
	// StillPlaying until one of the fleet not destroyed
	StillPlaying bool
	// GameEndReason if game ends then will be known reason of ending
	GameEndReason string
	// bot to emulate gaming
	bot *ai.Govern
}

// TODO Validate size of fleet and types of the ships

// NewGame creates game process
func NewGame(playerName string, playerFleet []*game.Ship) (*GameMaster, error) {
	var errPlayer error
	gp := &GameMaster{
		room:         game.NewWarRoom(),
		StillPlaying: true,
		bot:          ai.NewGovern(),
	}

	_, errPlayer = game.NewPlayer(playerName, playerFleet, gp.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	// TODO Add generator for fleet creation
	shipCoordinate := game.Coordinate{AxisX: 0, AxisY: 0}
	shipLocation := make([]game.Coordinate, 1)
	shipLocation[0] = shipCoordinate
	fleet := make([]*game.Ship, 0)
	fleet = append(fleet, &game.Ship{IsAlive: true, Location: shipLocation})
	_, errPlayer = game.NewPlayer(gp.bot.GetName(), fleet, gp.room)
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
	for {
		// AI did action
		isBotHit := nextPlayer.GunShoot(gp.bot.OpenFire())
		// Check if player has fleet
		gp.checkPlayerFleet(gp.room.GetOppositePlayer(gp.bot.GetName()))
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
