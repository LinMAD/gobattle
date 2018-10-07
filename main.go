package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/LinMAD/gobattle/pkg"
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"github.com/LinMAD/gobattle/pkg/render"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	newGame         *pkg.GameMaster
	isGameWithHuman bool
	playerName      string
)

// init game setup
func init() {
	var newGameErr error

	flag.StringVar(&playerName, "name", "MyPlayerName", "Players name")
	flag.BoolVar(&isGameWithHuman, "isHuman", false, "Is Game with human")
	flag.Parse()

	// Setup player, name, fleet
	playerFleet := generator.NewFleet()
	render.ShowBattleField(
		render.Screen{
			Title:       playerName + " it's your fleet" ,
			BattleField: generator.NewSeaField(playerFleet),
		},
		true,
	)
	settings := pkg.GameSettings{
		PlayerName:    playerName,
		PlayerFleet:   generator.NewFleet(),
		IsVersusHuman: isGameWithHuman,
		GameSpeed:     100,
	}

	newGame, newGameErr = pkg.NewGame(settings)
	if newGameErr != nil {
		log.Println(newGameErr.Error())
	}
}

// main, game loop starts here with condition of game type Human vs Bot or self play
func main() {
	seaPlan := generator.NewSeaField(nil)
	if !isGameWithHuman {
		for newGame.StillPlaying {
			// TODO Implement own AI\Bot to win the game
			newGame.ShootInCoordinate(game.Coordinate{AxisX: 1, AxisY: 1})
		}
	} else {
		var isHit bool
		reader := bufio.NewReader(os.Stdin)
		isNextCycle := false
		for newGame.StillPlaying {
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

			isHit = newGame.ShootInCoordinate(target)
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

	fmt.Println("--- GAME END ---")
	fmt.Printf("--- %s --- \n", newGame.GameEndReason)
}

// clearHumanInput handle input
func clearHumanInput(input string) int8 {
	re := regexp.MustCompile(`(\r?\n)|\s`)

	str := re.ReplaceAllString(input, "")
	strInt, _ := strconv.Atoi(str)

	return int8(strInt)
}
