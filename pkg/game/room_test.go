package game

import "testing"

func TestNewWarRoom(t *testing.T) {
	wr := NewWarRoom()
	if wr.players.Len() != 0 {
		t.Error("For new room expected empty players")
	}

	p1 := NewPlayer("P1")
	wr.AddPlayer(p1)
	if wr.players.Len() != 1 {
		t.Error("Expected to be added 1 player to room")
	}

	wr.AddPlayer(p1)
	if wr.players.Len() != 1 {
		t.Error("Expected 1 player in room, because was added same player")
	}

	p2 := NewPlayer("P2")
	wr.AddPlayer(p2)
	if wr.players.Len() != 2 {
		t.Error("Expected 2 player in room")
	}
}
