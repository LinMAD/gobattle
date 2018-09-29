package game

import "testing"

func TestNewPlayer(t *testing.T) {
	fleet := make([]Ship, 0)
	name := "Gopher"
	p, err := NewPlayer(name, fleet)

	if err == nil {
		t.Error("Expected to be error, player fleet empty")
	}

	fleet = append(fleet, Ship{})
	p, err = NewPlayer(name, fleet)

	if p.GetName() != name {
		t.Error("Player created incorrect, name not same -> " + name)
	}

	if len(p.GetFleet()) == 0 {
		t.Error("Expected empty list of ships for new player")
	}
}
