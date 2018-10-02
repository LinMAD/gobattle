package main

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"log"
)

var player *game.Player

// TODO Refactor markup when mechanism with "AI" interaction will be implemented

func init()  {
	// TODO Add component to return prepared dependencies with "AI"
	warRoom := game.NewWarRoom()

	// TODO Add generator for fleet creation
	shipCoordinate := game.Coordinate{AxisX: 0, AxisY: 0}
	shipLocation := make([]game.Coordinate, 1)
	shipLocation[0] = shipCoordinate
	fleet := make([]*game.Ship, 0)
	fleet = append(fleet, &game.Ship{IsAlive:  true, Location: shipLocation})

	player, errPlayer := game.NewPlayer("MyPlayerName", fleet, warRoom)
	if errPlayer != nil {
		log.Fatalln(errPlayer.Error())
	}

	warRoom.JoinPlayer(player)
}

func main() {
	//player.GunShoot(game.Coordinate{AxisX: 1, AxisY: 1})
}
