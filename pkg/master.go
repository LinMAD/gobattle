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
	gp := &GameMaster{
		room:         game.NewWarRoom(),
		StillPlaying: true,
		bot:          ai.NewGovern("Govern"),
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

		// Ok if bot win or miss then return back control to player
		if gp.StillPlaying == false || isBotHit == false {
			break
		}
	}

	return isFirstPlayerShot
}

// HandleHumanPlayer handle user input
func (gp *GameMaster) HandleHumanPlayer(playerName string) {
	var isHit bool

	seaPlan := generator.NewSeaField(nil)
	reader := bufio.NewReader(os.Stdin)
	isNextCycle := false

	for gp.StillPlaying {
		render.ShowBattleField(
			render.Screen{
				Title:       "Battle field of " + playerName,
				BattleField: seaPlan,
			},
			isNextCycle,
		)
		fmt.Println("Enter coordinate to fire.")

		target := game.Coordinate{}
		fmt.Print("Target X coordinate: ")
		xStr, _ := reader.ReadString('\n')
		target.AxisX = clearHumanInput(xStr)

		fmt.Print("Target Y coordinate: ")
		yStr, _ := reader.ReadString('\n')
		target.AxisY = clearHumanInput(yStr)

		isHit = gp.ShootInCoordinate(target)
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
