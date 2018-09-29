package game

import (
	"container/list"
	"fmt"
)

// WarRoomMediator
type WarRoomMediator interface {
	// AddPlayer to room
	AddPlayer()
}

// WarRoom
type WarRoom struct {
	players *list.List
}

// NewWarRoom
func NewWarRoom() *WarRoom {
	return &WarRoom{players: list.New()}
}

// AddPlayer to room with his fleet
func (room *WarRoom) AddPlayer(newPlayer string, fleet []Ship) error {
	isAdded := room.findPlayer(newPlayer)
	if isAdded != nil {
		return fmt.Errorf("player must be unique in room")
	}

	p, err := NewPlayer(newPlayer, fleet)
	if err != nil {
		return err
	}

	room.players.PushBack(p)

	return nil
}

// findPlayer in current room
func (room *WarRoom) findPlayer(playerName string) PlayerColleague {
	for p := room.players.Front(); p != nil; p = p.Next() {
		if p.Value.(PlayerColleague).GetName() == playerName {
			return p.Value.(PlayerColleague)
		}
	}

	return nil
}

// TODO Implement turn based Q for players
