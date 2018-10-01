package game

import (
	"container/list"
	"fmt"
)

// WarRoomMediator
type WarRoomMediator interface {
	// JoinPlayer to room
	JoinPlayer(p *Player) error
	// MakeTurn of active player
	MakeTurn(p *Player)
}

// WarRoom
type WarRoom struct {
	players *list.List
}

// NewWarRoom
func NewWarRoom() *WarRoom {
	return &WarRoom{players: list.New()}
}

// JoinPlayer to room with his fleet
func (room *WarRoom) JoinPlayer(p *Player) error {
	isAdded := room.findPlayer(p.name)
	if isAdded != nil {
		return fmt.Errorf("player must be unique in room")
	}

	room.players.PushBack(p)

	return nil
}

// findPlayer in current room
func (room *WarRoom) findPlayer(playerName string) *Player {
	for p := room.players.Front(); p != nil; p = p.Next() {
		if p.Value.(*Player).name == playerName {
			return p.Value.(*Player)
		}
	}

	return nil
}

// MakeTurn for player
func (room *WarRoom) MakeTurn(p *Player) {
	//target := p.lastFireCoordinate
	// TODO Get opponent ships and hit ships if target correct
	// TODO Return result of shooting, like miss, hit or kill

	// TODO Implement turn based Q for players
}
