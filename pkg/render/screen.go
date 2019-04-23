package render

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/game"
)

var createdScreens []Screen

func init()  {
	if createdScreens == nil {
		createdScreens = make([]Screen, 0)
	}
}

// Screen to show
type Screen struct {
	// Title of screen
	Title string
	// BattleField in action
	BattleField [][]string
}

// AddOrUpdate for rendering
func AddOrUpdate(s Screen)  {
	isNew := true
	for i, cs := range createdScreens {
		if cs.Title == s.Title {
			createdScreens[i] = s
			isNew = false
		}
	}

	if isNew {
		createdScreens = append(createdScreens, s)
	}
}

// ShowBattleField show battle field
func ShowBattleField(s Screen, clearScreen bool) {
	AddOrUpdate(s)

	screenBuffer := "\n"
	lenScreens := len(createdScreens)

	for i := 0; i < lenScreens; i++ {
		nI := i+1
		if lenScreens <= nI {
			break
		}

		/**
			Deeding on one or two fields will be assembled sections
			1. Title (Player name and field)
			2. Header
			3. Field rows (With left tile and coordinates)
			4. Footer tile
			5. Footer with coordinates
		 */

		if nI % 2 == 0 {
			screenBuffer += fmt.Sprintf(
				"\n%s               %s",
				createdScreens[i].buildTitle(),
				createdScreens[nI].buildTitle(),
			)
			screenBuffer += fmt.Sprintf(
				"\n%s                                    %s",
				createdScreens[i].buildHeader(),
				createdScreens[nI].buildHeader(),
			)

			for y := game.FSize - 1; y != 0; y-- {
				screenBuffer += fmt.Sprintf(
					"\n%s    %s",
					createdScreens[i].buildRow(y),
					createdScreens[nI].buildRow(y),
				)
			}

			screenBuffer += fmt.Sprintf(
				"\n%s       %s",
				createdScreens[i].buildFooterTile(),
				createdScreens[nI].buildFooterTile(),
			)
			screenBuffer += fmt.Sprintf(
				"\n  %s     %s",
				createdScreens[i].buildFooter(),
				createdScreens[nI].buildFooter(),
			)
		} else {
			screenBuffer += fmt.Sprintf("\n\n%s", createdScreens[i].buildTitle())
			screenBuffer += fmt.Sprintf("\n%s", createdScreens[i].buildHeader())

			for y := game.FSize - 1; y != 0; y-- {
				screenBuffer += fmt.Sprintf(
					"\n%s", createdScreens[i].buildRow(y))
			}

			screenBuffer += fmt.Sprintf("\n%s", createdScreens[i].buildFooterTile())
			screenBuffer += fmt.Sprintf("\n  %s", createdScreens[i].buildFooter())
			screenBuffer += fmt.Sprint("\n")
		}
	}

	if clearScreen {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Print(screenBuffer)
	screenBuffer = "" // clean
}

func (s *Screen) buildTitle() string {
	return s.Title
}

func (s *Screen) buildHeader() string {
	return "Y"
}

func (s *Screen) buildFooterTile() string {
	var tile string

	for x := 0; x < game.FSize; x++ {
		tile += "---"
	}

	return tile
}

func (s *Screen) buildFooter() string {
	footer := "X"

	for x := 0; x < game.FSize; x++ {
		footer += fmt.Sprintf(" %d ", x)
	}

	return footer
}

func (s *Screen) buildRow(y int) string {
	var row string

	row += fmt.Sprintf("%d| ", y)
	for x := 0; x < game.FSize; x++ {
		row += fmt.Sprintf(" %s ", s.BattleField[y][x])
	}

	return row
}
