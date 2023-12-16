package days

import (
	"bufio"
	"fmt"
	"os"
)

var prisms = [][]rune{}

func Day16() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day16Sample.txt"
	} else {
		fileName = "inputfiles/Day16.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var prismRow []rune
		for _, r := range scanner.Text() {
			prismRow = append(prismRow, r)
		}
		prisms = append(prisms, prismRow)
	}
	//Part 1
	tilesHit := make(map[[3]rune]interface{})
	energizedTiles := make(map[[2]int]interface{})
	energizedTiles = shootBeam(0, -1, 'R', energizedTiles, tilesHit)
	fmt.Println(len(energizedTiles))
	//Part 2
	var maxTiles int
	for col := range prisms[0] {
		energizedTiles := make(map[[2]int]interface{})
		tilesHit := make(map[[3]rune]interface{})
		energizedTiles = shootBeam(-1, col, 'D', energizedTiles, tilesHit)
		if len(energizedTiles) > maxTiles {
			maxTiles = len(energizedTiles)
		}
	}
	for col := range prisms[len(prisms)-1] {
		energizedTiles := make(map[[2]int]interface{})
		tilesHit := make(map[[3]rune]interface{})
		energizedTiles = shootBeam(len(prisms), col, 'U', energizedTiles, tilesHit)
		if len(energizedTiles) > maxTiles {
			maxTiles = len(energizedTiles)
		}
	}
	for row := 0; row < len(prisms); row++ {
		energizedTiles := make(map[[2]int]interface{})
		tilesHit := make(map[[3]rune]interface{})
		energizedTiles = shootBeam(row, -1, 'R', energizedTiles, tilesHit)
		if len(energizedTiles) > maxTiles {
			maxTiles = len(energizedTiles)
		}
		energizedTiles = make(map[[2]int]interface{})
		tilesHit = make(map[[3]rune]interface{})
		energizedTiles = shootBeam(row, len(prisms[row]), 'L', energizedTiles, tilesHit)
		if len(energizedTiles) > maxTiles {
			maxTiles = len(energizedTiles)
		}
	}
	fmt.Println(maxTiles)
}

func shootBeam(row int, col int, direction rune, energizedTiles map[[2]int]interface{}, tilesHit map[[3]rune]interface{}) map[[2]int]interface{} {
	switch direction {
	case 'R':
		col++
		if col < len(prisms[row]) {
			if _, ok := tilesHit[[3]rune{rune(row), rune(col), direction}]; !ok {
				tilesHit[[3]rune{rune(row), rune(col), direction}] = nil
				if _, ok := energizedTiles[[2]int{row, col}]; !ok {
					energizedTiles[[2]int{row, col}] = nil
				}
				switch prisms[row][col] {
				case '\\':
					shootBeam(row, col, 'D', energizedTiles, tilesHit)
				case '|':
					shootBeam(row, col, 'U', energizedTiles, tilesHit)
					shootBeam(row, col, 'D', energizedTiles, tilesHit)
				case '-':
					shootBeam(row, col, 'R', energizedTiles, tilesHit)
				case '/':
					shootBeam(row, col, 'U', energizedTiles, tilesHit)
				case '.':
					shootBeam(row, col, 'R', energizedTiles, tilesHit)
				}
			}
		}
	case 'D':
		row++
		if row < len(prisms) {
			if _, ok := tilesHit[[3]rune{rune(row), rune(col), direction}]; !ok {
				tilesHit[[3]rune{rune(row), rune(col), direction}] = nil
				if _, ok := energizedTiles[[2]int{row, col}]; !ok {
					energizedTiles[[2]int{row, col}] = nil
				}
				switch prisms[row][col] {
				case '\\':
					shootBeam(row, col, 'R', energizedTiles, tilesHit)
				case '|':
					shootBeam(row, col, 'D', energizedTiles, tilesHit)
				case '-':
					shootBeam(row, col, 'R', energizedTiles, tilesHit)
					shootBeam(row, col, 'L', energizedTiles, tilesHit)
				case '/':
					shootBeam(row, col, 'L', energizedTiles, tilesHit)
				case '.':
					shootBeam(row, col, 'D', energizedTiles, tilesHit)
				}
			}
		}
	case 'L':
		col--
		if col >= 0 {
			if _, ok := tilesHit[[3]rune{rune(row), rune(col), direction}]; !ok {
				tilesHit[[3]rune{rune(row), rune(col), direction}] = nil
				if _, ok := energizedTiles[[2]int{row, col}]; !ok {
					energizedTiles[[2]int{row, col}] = nil
				}
				switch prisms[row][col] {
				case '\\':
					shootBeam(row, col, 'U', energizedTiles, tilesHit)
				case '|':
					shootBeam(row, col, 'U', energizedTiles, tilesHit)
					shootBeam(row, col, 'D', energizedTiles, tilesHit)
				case '-':
					shootBeam(row, col, 'L', energizedTiles, tilesHit)
				case '/':
					shootBeam(row, col, 'D', energizedTiles, tilesHit)
				case '.':
					shootBeam(row, col, 'L', energizedTiles, tilesHit)
				}
			}
		}
	case 'U':
		row--
		if row >= 0 {
			if _, ok := tilesHit[[3]rune{rune(row), rune(col), direction}]; !ok {
				tilesHit[[3]rune{rune(row), rune(col), direction}] = nil
				if _, ok := energizedTiles[[2]int{row, col}]; !ok {
					energizedTiles[[2]int{row, col}] = nil
				}
				switch prisms[row][col] {
				case '\\':
					shootBeam(row, col, 'L', energizedTiles, tilesHit)
				case '|':
					shootBeam(row, col, 'U', energizedTiles, tilesHit)
				case '-':
					shootBeam(row, col, 'R', energizedTiles, tilesHit)
					shootBeam(row, col, 'L', energizedTiles, tilesHit)
				case '/':
					shootBeam(row, col, 'R', energizedTiles, tilesHit)
				case '.':
					shootBeam(row, col, 'U', energizedTiles, tilesHit)
				}
			}
		}
	}
	return energizedTiles
}
