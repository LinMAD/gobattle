package main

import (
	"github.com/LinMAD/gobattle/pkg"
	"github.com/LinMAD/gobattle/pkg/game"
	"log"
)

var newGame *pkg.GameMaster

func init() {
	var newGameErr error

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

// main represents game loop
func main() {
	for newGame.StillPlaying {
		// TODO Implement own AI\Bot to win the game
		fireTo := game.Coordinate{AxisX: 1, AxisY: 1}

		isHit := newGame.ShootInCoordinate(fireTo)
		log.Println("Did I shoot the ship? It's ", isHit)
	}

	log.Println("--- GAME END ---")
	log.Printf("--- %s --- \n", newGame.GameEndReason)
}
