package game

import (
	"container/list"
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

// AddPlayer to room
func (room *WarRoom) AddPlayer(newPlayer PlayerColleague) {
	isAdded := false

	// Check if player already added
	for player := room.players.Front(); player != nil; player = player.Next() {
		if player.Value.(PlayerColleague).GetName() == newPlayer.GetName() {
			isAdded = true
		}
	}

	if !isAdded {
		room.players.PushBack(newPlayer)
	}
}

// TODO Implement turn based Q for players
