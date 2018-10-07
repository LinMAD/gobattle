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
	playerName string
)

// init game setup
func init() {
	var newGameErr error

	flag.StringVar(&playerName, "name", "MyPlayerName", "Players name")
	flag.BoolVar(&isGameWithHuman, "isHuman", false, "Is Game with human")
	flag.Parse()

	// Setup player, name, fleet
	newGame, newGameErr = pkg.NewGame(playerName, generator.NewFleet())
	if newGameErr != nil {
		log.Println(newGameErr.Error())
	}
}

// main, game loop starts here with condition of game type Human vs Bot or self play
func main() {
	if !isGameWithHuman {
		for newGame.StillPlaying {
			// TODO Implement own AI\Bot to win the game
			newGame.ShootInCoordinate(game.Coordinate{AxisX: 1, AxisY: 1})
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		seaPlan := generator.NewSeaField(nil)
		for newGame.StillPlaying {
			render.ShowBattleField("Battle field of " + playerName, seaPlan)

			target := game.Coordinate{}
			fmt.Println("Enter coordinate to fire.")

			fmt.Print("Target X coordinate: ")
			xStr, _ := reader.ReadString('\n')
			target.AxisX = clearHumanInput(xStr)

			fmt.Print("Target Y coordinate: ")
			yStr, _ := reader.ReadString('\n')
			target.AxisY = clearHumanInput(yStr)

			isHit := newGame.ShootInCoordinate(target)
			if isHit {
				seaPlan[target.AxisY][target.AxisX] = 5
			} else {
				seaPlan[target.AxisY][target.AxisX] = game.FShot
			}

			fmt.Printf("Did I shoot the ship? %v \n\n", isHit)
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
