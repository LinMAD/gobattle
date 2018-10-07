package render

import (
	"fmt"
	"github.com/LinMAD/gobattle/pkg/game"
)

// ShowBattleField show battle field
func ShowBattleField(title string, battleField [][]int8) {
	fmt.Print("\033[H\033[2J\n")
	fmt.Println(title)
	fmt.Println("\nY")

	underLine := ""
	for y := 9; y >= 0; y-- {
		fmt.Printf("%d| ", y)

		underLine += "---"
		for x := 0; x < 10; x++ {
			if battleField[y][x] == game.FShot {
				fmt.Printf("%d ", battleField[y][x])
				continue
			}
			fmt.Printf(" %d ", battleField[y][x])
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
