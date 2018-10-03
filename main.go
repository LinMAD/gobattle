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
	myFleet := make([]*game.Ship, 0)

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
		log.Println("Did i shoot the ship: ", isHit)
		// TODO Verify my fleet, if I miss my shot then I can have casualties
	}
}
