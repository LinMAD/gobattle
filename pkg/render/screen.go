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
	if clearScreen {
		fmt.Print("\033[H\033[2J")
	}

	fmt.Println("\n", s.Title)
	fmt.Println("\nY")

	underLine := ""
	for y := 9; y >= 0; y-- {
		fmt.Printf("%d| ", y)

		underLine += "---"
		for x := 0; x < 10; x++ {
			fmt.Printf(" %s ", s.BattleField[y][x])
		}
		fmt.Println()

		if y == 0 {
			fmt.Printf("%s--\n   ", underLine)
			for i := 0; i <= 9; i++ {
				fmt.Printf(" %d ", i)
			}
			fmt.Print("X")
			fmt.Println()
		}
	}
}
