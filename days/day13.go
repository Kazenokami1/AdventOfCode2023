package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day13() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day13Sample.txt"
	} else {
		fileName = "inputfiles/Day13.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inputByRow [][]string
	var inputByColumn [][]string
	var rows []string
	for scanner.Scan() {
		if scanner.Text() != "" {
			rows = append(rows, scanner.Text())
		} else {
			inputByRow = append(inputByRow, rows)
			rows = []string{}
		}
	}
	inputByRow = append(inputByRow, rows)
	for _, input := range inputByRow {
		var cols []string
		for i := 0; i < len(input[0]); i++ {
			var col string
			for j := 0; j < len(input); j++ {
				col += string(input[j][i])
			}
			cols = append(cols, col)
		}
		inputByColumn = append(inputByColumn, cols)
	}
	day13part1(inputByRow, inputByColumn)
	day13part2(inputByRow, inputByColumn)
}

func day13part1(inputByRow [][]string, inputByColumn [][]string) {
	var summary int
	for _, input := range inputByRow {
		var reflectionFound bool
		var differences int
		var i int
		for !reflectionFound {
			differences = compareStrings(input, i, i+1, math.Min(float64(i), float64(len(input)-2-i)))
			if differences == 0 {
				summary += (i + 1) * 100
				reflectionFound = true
			} else {
				i++
				if i > len(input)-2 {
					break
				}
			}
		}
	}
	for _, input := range inputByColumn {
		var reflectionFound bool
		var differences int
		var i int
		for !reflectionFound {
			differences = compareStrings(input, i, i+1, math.Min(float64(i), float64(len(input)-2-i)))
			if differences == 0 {
				summary += (i + 1)
				reflectionFound = true
			} else {
				i++
				if i > len(input)-2 {
					break
				}
			}
		}
	}
	fmt.Println(summary)
}

func day13part2(inputByRow [][]string, inputByColumn [][]string) {
	var summary int
	for _, input := range inputByRow {
		var reflectionFound bool
		var differences int
		var i int
		for !reflectionFound {
			differences = compareStrings(input, i, i+1, math.Min(float64(i), float64(len(input)-2-i)))
			if differences == 1 {
				summary += (i + 1) * 100
				reflectionFound = true
			} else {
				i++
				if i > len(input)-2 {
					break
				}
			}
		}
	}
	for _, input := range inputByColumn {
		var reflectionFound bool
		var differences int
		var i int
		for !reflectionFound {
			differences = compareStrings(input, i, i+1, math.Min(float64(i), float64(len(input)-2-i)))
			if differences == 1 {
				summary += (i + 1)
				reflectionFound = true
			} else {
				i++
				if i > len(input)-2 {
					break
				}
			}
		}
	}
	fmt.Println(summary)
}

func compareStrings(input []string, first int, second int, recursions float64) int {
	var differences int
	if recursions > 0 {
		differences += compareStrings(input, first-1, second+1, recursions-1)
	}
	if differences > 1 {
		return differences
	} else {
		for i := 0; i < len(input[first]); i++ {
			if input[first][i] != input[second][i] {
				differences++
			}
		}
	}
	return differences
}
