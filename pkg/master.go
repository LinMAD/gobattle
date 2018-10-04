package pkg

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"log"
)

const botName = "Govern"

// GameMaster encapsulates game iterations and follows rules
type GameMaster struct {
	// room with players
	room *game.WarRoom
	// StillPlaying until one of the fleet not destroyed
	StillPlaying bool
}

// TODO Validate size of fleet and types of the ships
// TODO Add game end reason, who win
// TODO Cover master with unit tests

// NewGame creates game process
func NewGame(playerName string, playerFleet []*game.Ship) (*GameMaster, error) {
	var errPlayer error
	gp := &GameMaster{
		room:         game.NewWarRoom(),
		StillPlaying: true,
	}

	_, errPlayer = game.NewPlayer(playerName, playerFleet, gp.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	// TODO Generate player for game AI
	// TODO Add generator for fleet creation
	shipCoordinate := game.Coordinate{AxisX: 0, AxisY: 0}
	shipLocation := make([]game.Coordinate, 1)
	shipLocation[0] = shipCoordinate
	fleet := make([]*game.Ship, 0)
	fleet = append(fleet, &game.Ship{IsAlive: true, Location: shipLocation})
	_, errPlayer = game.NewPlayer(botName, fleet, gp.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	return gp, nil
}

// ShootInCoordinate
func (gp *GameMaster) ShootInCoordinate(target game.Coordinate) bool {
	var isFirstPlayerShot bool

	// TODO Check if players still have a fleet, if not stop game

	// Here can be tricky
	// if next turn not a build in bot then
	// return only if shot hit a ship
	nextPlayer := gp.room.GetActivePlayer()
	if nextPlayer.GetName() != botName {
		isFirstPlayerShot = nextPlayer.GunShoot(target)
		if isFirstPlayerShot {
			return isFirstPlayerShot
		}
	}

	// So first player loosed initiative
	// Now bot will try shot all ships while first player waits
	// TODO Do that go routine
	var botHit bool // TODO remove that stub
	for {
		// TODO Give AI to choose where to shoot
		// AI did action
		log.Println("BOT SHOT")
		if botHit == false {
			return isFirstPlayerShot
		}
		botHit = false
	}
}
