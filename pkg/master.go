package pkg

import (
	"bufio"
	"fmt"
	"github.com/LinMAD/gobattle/pkg/ai"
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"github.com/LinMAD/gobattle/pkg/render"
	"os"
	"regexp"
	"strconv"
)

// GameSettings for new game
type GameSettings struct {
	// PlayerName on display
	PlayerName string
	// PlayerFleet collection of ships
	PlayerFleet []*game.Ship
	// IsVersusHuman flag to show screens
	IsVersusHuman bool
}

// GameMaster acts mostly like proxy to follow the game rules, encapsulates mechanics
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

// NewGame creates game to play
func NewGame(settings GameSettings) (*GameMaster, error) {
	var errPlayer error
	gm := &GameMaster{
		room:         game.NewWarRoom(),
		StillPlaying: true,
		bot:          ai.NewGovern("Govern"),
	}

	_, errPlayer = game.NewPlayer(settings.PlayerName, settings.PlayerFleet, gm.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	_, errPlayer = game.NewPlayer(gm.bot.GetName(), generator.NewFleet(), gm.room)
	if errPlayer != nil {
		return nil, errPlayer
	}

	return gm, nil
}

// ShootInCoordinate provide points hit ship if it's there
func (gm *GameMaster) ShootInCoordinate(target game.Coordinate) bool {
	var isFirstPlayerShot bool

	// Here can be tricky
	// if next turn not a build in bot then
	// return only if shot hit a ship
	nextPlayer := gm.room.GetActivePlayer()
	if nextPlayer.GetName() != gm.bot.GetName() {
		isFirstPlayerShot = nextPlayer.GunShoot(target)
		gm.checkPlayerFleet(gm.room.GetOppositePlayer(nextPlayer.GetName()))
		if isFirstPlayerShot {
			return isFirstPlayerShot
		}
	}

	// So first player loosed initiative
	// Now bot will try shot all ships while first player waits
	nextPlayer = gm.room.GetActivePlayer()
	oppositePlayer := gm.room.GetOppositePlayer(nextPlayer.GetName())
	for {
		// AI did action
		targetToHit := gm.bot.OpenFire()
		isBotHit := nextPlayer.GunShoot(targetToHit)
		// Collect result for bot
		gm.bot.CollectResultOfShot(targetToHit, isBotHit)
		render.AddOrUpdate(render.Screen{Title: "Battlefield of " + gm.bot.GetName(), BattleField: gm.bot.GetSeaPlan()})
		// Check if bot win the game
		gm.checkPlayerFleet(oppositePlayer)

		// Ok if bot win or miss then return back control to player
		if gm.StillPlaying == false || isBotHit == false {
			break
		}
	}

	return isFirstPlayerShot
}

// HandleHumanPlayer handle user input
func (gm *GameMaster) HandleHumanPlayer(playerName string) {
	var isHit bool

	seaPlan := generator.NewSeaField(nil)
	reader := bufio.NewReader(os.Stdin)
	isNextCycle := false

	for gm.StillPlaying {
		render.ShowBattleField(
			render.Screen{
				Title:       "Battlefield of " + playerName,
				BattleField: seaPlan,
			},
			isNextCycle,
		)

		fmt.Println("\n\nEnter coordinate to fire:")
		target := game.Coordinate{}
		isHandled := map[string]bool{"Y": false, "X": false}

		for {
			// Handle Y input
			if isHandled["Y"] == false {
				fmt.Print("Target Y coordinate: ")
				yStr, _ := reader.ReadString('\n')
				target.AxisY = clearHumanInput(yStr)

				if target.AxisY >= int8(game.FSize) {
					fmt.Println("Incorrect Y coordinate, try one more time")
					target.AxisY = 0
					continue
				}

				isHandled["Y"] = true
			}

			// Handle X input
			if isHandled["X"] == false {
				fmt.Print("Target X coordinate: ")
				xStr, _ := reader.ReadString('\n')
				target.AxisX = clearHumanInput(xStr)

				if target.AxisX >= int8(game.FSize) {
					fmt.Println("Incorrect X coordinate, try one more time")
					target.AxisX = 0
					continue
				}

				isHandled["X"] = true
			}

			break
		}

		isHit = gm.ShootInCoordinate(target)
		marker := seaPlan[target.AxisY][target.AxisX]
		if marker != game.GunMis && marker != game.GunHit {
			if isHit {
				seaPlan[target.AxisY][target.AxisX] = game.GunHit
			} else {
				seaPlan[target.AxisY][target.AxisX] = game.GunMis
			}
		}
		isNextCycle = true
	}
}

// clearHumanInput handle input
func clearHumanInput(input string) int8 {
	re := regexp.MustCompile(`(\r?\n)|\s`)

	str := re.ReplaceAllString(input, "")
	strInt, _ := strconv.Atoi(str)

	return int8(strInt)
}

// checkPlayerFleet check if game ended
func (gm *GameMaster) checkPlayerFleet(p *game.Player) {
	fleetSize := len(p.GetFleet())

	for _, s := range p.GetFleet() {
		if s.IsAlive == false {
			fleetSize--
		}
	}

	if fleetSize == 0 {
		gm.StillPlaying = false
		gm.GameEndReason = fmt.Sprintf("%s defeated, whole fleet destroyed", p.GetName())

		return
	}
}
