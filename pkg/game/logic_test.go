package game

import "testing"

func TestNewBattleGround(t *testing.T) {
	bg := NewBattleGround(5,5)
	if bg.MapSettings.AxisX != 5 || bg.MapSettings.AxisY != 5 {
		t.Error("New battle ground size not same as given")
	}
}
