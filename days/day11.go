package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day11() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day11Sample.txt"
	} else {
		fileName = "inputfiles/Day11.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var space []string
	columnsWithoutGalaxies := make(map[int]interface{})
	rowsWithoutGalaxies := make(map[int]interface{})
	for scanner.Scan() {
		space = append(space, scanner.Text())
	}
	for i := 0; i < len(space[0]); i++ {
		columnsWithoutGalaxies[i] = nil
	}
	for i := 0; i < len(space); i++ {
		rowsWithoutGalaxies[i] = nil
	}
	var galaxyMapPart1 [][]int
	var galaxyMapPart2 [][]int
	for row, spaceRow := range space {
		for col, symbol := range spaceRow {
			if symbol == '#' {
				galaxyMapPart1 = append(galaxyMapPart1, []int{row, col})
				galaxyMapPart2 = append(galaxyMapPart2, []int{row, col})
				delete(rowsWithoutGalaxies, row)
				delete(columnsWithoutGalaxies, col)
			}
		}
	}
	for i := range galaxyMapPart1 {
		var newRows int
		for row := range rowsWithoutGalaxies {
			if galaxyMapPart1[i][0] > row {
				newRows++
			}
		}
		galaxyMapPart1[i][0] += newRows
		galaxyMapPart2[i][0] += newRows * 999999
	}
	for i := range galaxyMapPart1 {
		var newCols int
		for col := range columnsWithoutGalaxies {
			if galaxyMapPart1[i][1] > col {
				newCols++
			}
		}
		galaxyMapPart1[i][1] += newCols
		galaxyMapPart2[i][1] += newCols * 999999
	}
	var shortestDistancesPart1 float64
	var shortestDistancesPart2 float64
	for i := 0; i < len(galaxyMapPart1)-1; i++ {
		for j := i + 1; j < len(galaxyMapPart1); j++ {
			shortestDistancesPart1 += math.Abs(float64(galaxyMapPart1[i][0])-float64(galaxyMapPart1[j][0])) + math.Abs(float64(galaxyMapPart1[i][1])-float64(galaxyMapPart1[j][1]))
			shortestDistancesPart2 += math.Abs(float64(galaxyMapPart2[i][0])-float64(galaxyMapPart2[j][0])) + math.Abs(float64(galaxyMapPart2[i][1])-float64(galaxyMapPart2[j][1]))
		}
	}
	fmt.Println(int(shortestDistancesPart1))
	fmt.Println(int(shortestDistancesPart2))
}
