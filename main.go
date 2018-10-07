package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/LinMAD/gobattle/pkg"
	"github.com/LinMAD/gobattle/pkg/game"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	newGame         *pkg.GameMaster
	isGameWithHuman bool
)

// init game setup
func init() {
	var newGameErr error

	flag.BoolVar(&isGameWithHuman, "Human game", false, "Is Game with human")
	flag.Parse()

	// Generate your own fleet
	// TODO Add generator
	shipCoordinate := game.Coordinate{AxisX: 0, AxisY: 0}
	shipLocation := make([]game.Coordinate, 1)
	shipLocation[0] = shipCoordinate
	myFleet := make([]*game.Ship, 0)
	myFleet = append(myFleet, &game.Ship{IsAlive: true, Location: shipLocation})

	newGame, newGameErr = pkg.NewGame("MyPlayerName", myFleet)
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
		for newGame.StillPlaying {
			target := game.Coordinate{}
			fmt.Println("Enter coordinate to fire.")

			fmt.Print("Target X coordinate: ")
			xStr, _ := reader.ReadString('\n')
			target.AxisX = clearHumanInput(xStr)

			fmt.Print("Target Y coordinate: ")
			yStr, _ := reader.ReadString('\n')
			target.AxisY = clearHumanInput(yStr)

			isHit := newGame.ShootInCoordinate(target)
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
