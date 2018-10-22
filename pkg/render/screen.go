package render

import (
	"fmt"
)

// Screen to show
type Screen struct {
	// Title of screen
	Title string
	// BattleField in action
	BattleField [][]string
}

// ShowBattleField show battle field
func ShowBattleField(s Screen, clearScreen bool) {
	var preparedScreen, underLine string

	preparedScreen += fmt.Sprintln("\n", s.Title)
	preparedScreen += fmt.Sprintln("\nY")

	for y := 9; y >= 0; y-- {
		preparedScreen += fmt.Sprintf("%d| ", y)

		underLine += "---"
		for x := 0; x < 10; x++ {
			preparedScreen += fmt.Sprintf(" %s ", s.BattleField[y][x])
		}
		preparedScreen += fmt.Sprintln()

		if y == 0 {
			preparedScreen += fmt.Sprintf("%s--\n   ", underLine)
			for i := 0; i <= 9; i++ {
				preparedScreen += fmt.Sprintf(" %d ", i)
			}
			preparedScreen += fmt.Sprint("X")
			preparedScreen += fmt.Sprintln()
		}
	}

	if clearScreen {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Print(preparedScreen)
}
