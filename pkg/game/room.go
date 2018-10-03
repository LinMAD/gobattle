package game

import (
	"fmt"
)

// WarRoomMediator
type WarRoomMediator interface {
	// JoinPlayer to room
	JoinPlayer(p *Player) error
	// MakeTurn for player and return if he succeed
	MakeTurn(p *Player) bool
	// GetActivePlayer returns player who must make turn
	GetActivePlayer() *Player
}

// WarRoom
type WarRoom struct {
	players []*Player
}

// NewWarRoom
func NewWarRoom() *WarRoom {
	return &WarRoom{players: make([]*Player, 0)}
}

// JoinPlayer to room with his fleet
func (room *WarRoom) JoinPlayer(newPlayer *Player) error {
	if len(room.players) == 2 {
		return fmt.Errorf("in room can be only 2 players")
	}

	for _, inPlayerRoom := range room.players {
		if inPlayerRoom.name == newPlayer.name {
			return fmt.Errorf("player must be unique in room")
		}
	}

	room.players = append(room.players, newPlayer)

	return nil
}

// getOppositePlayer in room
func (room *WarRoom) getOppositePlayer(playerName string) *Player {
	for _, p := range room.players {
		if p.name != playerName {
			return p
		}
	}

	return nil
}

// MakeTurn for player and return if he succeed
func (room *WarRoom) MakeTurn(p *Player) bool {
	var isHit bool
	// is ship was damaged during firing in targeted coordinates
	targetCoordinate := p.lastFireCoordinate

	oppositePlayer := room.getOppositePlayer(p.name)

	// Go throw all ships and try hit
	for _, ship := range oppositePlayer.GetFleet() {
		isHit = ship.isHit(targetCoordinate)
	}

	return isHit
}

// GetActivePlayer he must make turn
func (room *WarRoom) GetActivePlayer() *Player {
	for _, p := range room.players {
		if p.isActive {
			return p
		}
	}

	// That should not happen, only if one player in room
	return nil
}
