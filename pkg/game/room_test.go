package game

import "testing"

func TestNewWarRoom(t *testing.T) {
	wr := NewWarRoom()
	if wr.players.Len() != 0 {
		t.Error("For new room expected empty players")
	}

	p1Fleet := make([]Ship, 1)
	p1Fleet[0].Location = append(p1Fleet[0].Location, Coordinate{1, 1})
	p1, err := NewPlayer("P1", p1Fleet, wr)
	if err != nil {
		t.Error(err.Error())
	}

	addErr := wr.JoinPlayer(p1)
	if wr.players.Len() != 1 {
		t.Error("Expected to be added 1 player to room")
	}

	addErr = wr.JoinPlayer(p1)
	if addErr == nil {
		t.Error(addErr)
	}
	if wr.players.Len() != 1 {
		t.Error("Expected 1 player in room, because was added same player")
	}

	p2Fleet := make([]Ship, 1)
	p2Fleet[0].Location = append(p1Fleet[0].Location, Coordinate{1, 1})
	_, addErr = NewPlayer("P2", p2Fleet, wr)
	if addErr != nil {
		t.Error(err.Error())
	}
	if wr.players.Len() != 2 {
		t.Error("Expected 2 player in room")
	}
}
