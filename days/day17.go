package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day17() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day17Sample.txt"
	} else {
		fileName = "inputfiles/Day17.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	gridSquares := make(map[[2]int]grid)
	var row int
	for scanner.Scan() {
		for col, heat := range scanner.Text() {
			gridSquares[[2]int{row, col}] = grid{heatValue: int(heat - '0'), heatLossFromStart: math.MaxInt}
		}
	}
	for pos, square := range gridSquares {
		for i := -1; i <= 1; i += 2 {
			if grid, ok := gridSquares[[2]int{pos[0] + i, pos[1]}]; ok {
				square.addNeighbor(&grid)
			}
			if grid, ok := gridSquares[[2]int{pos[0], pos[1] + i}]; ok {
				square.addNeighbor(&grid)
			}
		}
	}
	fmt.Println()
}
