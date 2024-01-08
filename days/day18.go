package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day18() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day18Sample.txt"
	} else {
		fileName = "inputfiles/Day18.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var digSteps [][]string
	for scanner.Scan() {
		digSteps = append(digSteps, strings.Split(scanner.Text(), " "))
	}
	var lagoonPart1 [][2]int
	var part1row int
	var part1col int
	var part1Perimeter int
	var lagoonPart2 [][2]int64
	var part2row int64
	var part2col int64
	var part2Perimeter int64
	for _, step := range digSteps {
		move, _ := strconv.Atoi(step[1])
		part1Perimeter += move
		switch step[0] {
		case "U":
			part1row -= move
		case "D":
			part1row += move
		case "R":
			part1col += move
		case "L":
			part1col -= move
		}
		lagoonPart1 = append(lagoonPart1, [2]int{part1col, part1row})
		part2string := step[2][2 : len(step[2])-1]
		movePart2, _ := strconv.ParseInt(part2string[0:len(part2string)-1], 16, 64)
		part2Perimeter += movePart2
		switch part2string[len(part2string)-1] {
		case '0':
			part2col += movePart2
		case '1':
			part2row += movePart2
		case '2':
			part2col -= movePart2
		case '3':
			part2row -= movePart2
		}
		lagoonPart2 = append(lagoonPart2, [2]int64{part2col, part2row})
	}
	var part1Area int
	var part2Area int64
	for i := 0; i < len(lagoonPart1); i++ {
		if i == len(lagoonPart1)-1 {
			part1Area += lagoonPart1[i][0] * lagoonPart1[0][1]
			part2Area += lagoonPart2[i][0] * lagoonPart2[0][1]
		} else {
			part1Area += lagoonPart1[i][0] * lagoonPart1[i+1][1]
			part2Area += lagoonPart2[i][0] * lagoonPart2[i+1][1]

		}
	}
	for i := 0; i < len(lagoonPart1); i++ {
		if i == len(lagoonPart1)-1 {
			part1Area -= lagoonPart1[i][1] * lagoonPart1[0][0]
			part2Area -= lagoonPart2[i][1] * lagoonPart2[0][0]
		} else {
			part1Area -= lagoonPart1[i][1] * lagoonPart1[i+1][0]
			part2Area -= lagoonPart2[i][1] * lagoonPart2[i+1][0]
		}
	}
	totalPointsEnclosed := float64(part1Area)*0.5 + float64(part1Perimeter)*0.5 + 1
	fmt.Println(totalPointsEnclosed)
	totalPointsEnclosedPart2 := float64(part2Area)*0.5 + float64(part2Perimeter)*0.5 + 1
	fmt.Println(totalPointsEnclosedPart2)
}
