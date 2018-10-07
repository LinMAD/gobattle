package game

import (
	"testing"
)

func TestNewWarRoom(t *testing.T) {
	wr := NewWarRoom()
	if len(wr.players) != 0 {
		t.Error("For new room expected empty players")
	}

	p1Fleet := helpCreateFleet(helpCreateShip(Coordinate{1, 1}))
	p1, err := NewPlayer("P1", p1Fleet, wr)
	if err != nil {
		t.Error(err.Error())
	}

	addErr := wr.JoinPlayer(p1)
	if len(wr.players) != 1 {
		t.Error("Expected to be added 1 player to room")
	}

	addErr = wr.JoinPlayer(p1)
	if addErr == nil {
		t.Error(addErr)
	}
	if len(wr.players) != 1 {
		t.Error("Expected 1 player in room, because was added same player")
	}

	p2Fleet := helpCreateFleet(helpCreateShip(Coordinate{1, 1}))
	_, addErr = NewPlayer("P2", p2Fleet, wr)
	if addErr != nil {
		t.Error(err.Error())
	}
	if len(wr.players) != 2 {
		t.Error("Expected 2 player in room")
	}

	_, addErr = NewPlayer("P3", p2Fleet, wr)
	if addErr != nil {
		t.Error("Expected error, in room can be only 2 players")
	}
}

func TestWarRoom_MakeTurn(t *testing.T) {
	wr, err := helpCreateRoom()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if wr.GetActivePlayer() == nil {
		t.Error("Expected one player be active")
	}

	p1 := wr.GetOppositePlayer("P2")
	if !p1.GunShoot(Coordinate{1, 1}) {
		t.Error("Expected to be damaged ship")
	}

	isDead := false
	p2 := wr.GetOppositePlayer("P1")
	for _, p2s := range p2.GetFleet() {
		if p2s.IsAlive == false {
			isDead = true
		}
	}

	if isDead == false {
		t.Error("One of the ships must be dead")
	}
}

func helpCreateRoom() (*WarRoom, error) {
	wr := NewWarRoom()

	_, p1Err := NewPlayer("P1", helpCreateFleet(helpCreateShip(Coordinate{1, 1})), wr)
	if p1Err != nil {
		return nil, p1Err
	}
	_, p2Err := NewPlayer(
		"P2",
		helpCreateFleet(
			helpCreateShip(Coordinate{5, 5}, Coordinate{5, 4}),
			helpCreateShip(Coordinate{3, 5}, Coordinate{3, 4}),
			helpCreateShip(Coordinate{1, 1}),
		),
		wr,
	)
	if p2Err != nil {
		return nil, p2Err
	}

	return wr, nil
}
