package game

import "testing"

func TestNewWarRoom(t *testing.T) {
	wr := NewWarRoom()
	if wr.players.Len() != 0 {
		t.Error("For new room expected empty players")
	}

	p1Fleet := make([]Ship, 0)
	p1 := "P1"
	addErr := wr.AddPlayer(p1, p1Fleet)
	if addErr == nil {
		t.Error("New player with empty fleet, expected error")
	}

	p1Fleet = append(p1Fleet, Ship{})
	p1Fleet[0].Location = append(p1Fleet[0].Location, Coordinate{1,1})
	addErr = wr.AddPlayer(p1, p1Fleet)
	if wr.players.Len() != 1 {
		t.Error("Expected to be added 1 player to room")
	}

	addErr = wr.AddPlayer(p1, p1Fleet)
	if addErr == nil {
		t.Error(addErr)
	}
	if wr.players.Len() != 1 {
		t.Error("Expected 1 player in room, because was added same player")
	}

	p2Fleet := make([]Ship, 1)
	p2Fleet[0].Location = append(p1Fleet[0].Location, Coordinate{1,1})
	p2 := "P2"
	wr.AddPlayer(p2, p2Fleet)
	if wr.players.Len() != 2 {
		t.Error("Expected 2 player in room")
	}
}
