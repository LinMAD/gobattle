package game

import "testing"

func TestNewPlayer(t *testing.T) {
	name := "Gopher"
	p := NewPlayer(name)

	if p.GetName() != name {
		t.Error("Player created incorrect, name not same -> " + name)
	}

	if len(p.GetShips()) != 0 {
		t.Error("Expected empty list of ships for new player")
	}
}
