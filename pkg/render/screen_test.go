package render

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/game"
	"github.com/LinMAD/gobattle/pkg/generator"
	"testing"
)

var ts * Screen

func init()  {
	ts = new(Screen)
	ts.Title = "Test title of screen"
	ts.BattleField = generator.NewSeaField(nil)
}

func TestTitleRender(t *testing.T) {
	if ts.Title != ts.buildTitle() {
		t.Fail()
	}
}

func TestHeaderRender(t *testing.T) {
	if ts.buildHeader() != "Y" {
		t.Fail()
	}
}

func TestRowRendering(t *testing.T) {
	var field string
	for y := game.FSize - 1; y >= 0; y-- {
		field += ts.buildRow(y)
	}

	if len(field) == 0 {
		t.Fail()
	}
}

// TestOneField debug render
func TestOneField(t *testing.T) {
	field := fmt.Sprintln(ts.buildTitle())
	field += fmt.Sprintln(ts.buildHeader())

	for y := game.FSize - 1; y >= 0; y-- {
		field += fmt.Sprintf("%s\n", ts.buildRow(y))
	}

	field += ts.buildFooter()
}
