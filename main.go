package main

import (
	"flag"
	"fmt"
	"github.com/LinMAD/gobattle/pkg"
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"github.com/LinMAD/gobattle/pkg/render"
	"log"
	"time"
)

var (
	gameMaster      *pkg.GameMaster
	gameSpeed       int
	isGameWithHuman bool
	playerName      string
)

// init game setup
func init() {
	var newGameErr error

	flag.StringVar(&playerName, "name", "Player", "Players name")
	flag.BoolVar(&isGameWithHuman, "isHuman", false, "Is Game with human")
	flag.IntVar(&gameSpeed, "gameSpeed", 100, "Game speed (for bots only)")
	flag.Parse()

	// Setup player, name, fleet
	playerFleet := generator.NewFleet()
	settings := pkg.GameSettings{
		PlayerName:    playerName,
		PlayerFleet:   playerFleet,
		IsVersusHuman: isGameWithHuman,
		GameSpeed:     time.Duration(gameSpeed),
	}

	gameMaster, newGameErr = pkg.NewGame(settings)
	if newGameErr != nil {
		log.Println(newGameErr.Error())
	}

	render.ShowBattleField(
		render.Screen{
			Title:       playerName + " it's your fleet",
			BattleField: generator.NewSeaField(playerFleet),
		},
		true,
	)
}

// main, game loop starts here with condition of game type Human vs Bot or self play
func main() {
	if isGameWithHuman {
		gameMaster.HandleHumanPlayer(playerName)
	} else {
		seaPlan := generator.NewSeaField(nil)

		for gameMaster.StillPlaying {

			// TODO Implement own AI\Bot to win the game

			// Set target to shoot foe ship
			target := game.Coordinate{
				AxisX: generator.RandomNum(0, game.FSize-1),
				AxisY: generator.RandomNum(0, game.FSize-1),
			}

			// Make move with given coordinates and record result
			if gameMaster.ShootInCoordinate(target) {
				seaPlan[target.AxisY][target.AxisX] = game.GunHit
			} else {
				seaPlan[target.AxisY][target.AxisX] = game.GunMis
			}

			// Show in screen enemy field with shot results
			render.ShowBattleField(
				render.Screen{
					Title:       "Battle field of " + playerName,
					BattleField: seaPlan,
				},
				true,
			)
		}
	}

	fmt.Println("--- GAME END ---")
	fmt.Printf("--- %s --- \n", gameMaster.GameEndReason)
}
