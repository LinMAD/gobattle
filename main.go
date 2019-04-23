package main

import (
	"flag"
	"fmt"
	"github.com/LinMAD/gobattle/pkg"
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"github.com/LinMAD/gobattle/pkg/render"
	"log"
)

var (
	playerName  string
	playerFleet []*game.Ship
)

// init game setup
func init() {
	flag.StringVar(&playerName, "name", "Player", "Players name")
	flag.Parse()

	// Setup player, name, fleet
	playerFleet = generator.NewFleet()
	render.AddOrUpdate(
		render.Screen{
			Title:       playerName + " it's your fleet",
			BattleField: generator.NewSeaField(playerFleet),
		},
	)
}

func main() {
	// create new game
	gameMaster, newGameErr := pkg.NewGame(
		pkg.GameSettings{PlayerName: playerName, PlayerFleet: playerFleet},
	)
	if newGameErr != nil {
		log.Fatalln(newGameErr)
	}

	// Handle game of Human vs Bot
	gameMaster.HandleHumanPlayer(playerName)

	fmt.Printf("\n\n--- GAME END ---")
	fmt.Printf("--- %s --- \n", gameMaster.GameEndReason)
}
