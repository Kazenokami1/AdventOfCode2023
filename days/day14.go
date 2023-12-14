package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirMap = make(map[int]string)

// 100 is the number of rows in my input, needed so I can use it as the key for a map.
// Adjust 100 to be the number of rows in your input (10 if you're running my sample data)
const rockLength = 100

func init() {
	dirMap[0] = "N"
	dirMap[1] = "W"
	dirMap[2] = "S"
	dirMap[3] = "E"
}

func Day14() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day14Sample.txt"
	} else {
		fileName = "inputfiles/Day14.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var rocks []string
	for scanner.Scan() {
		rocks = append(rocks, scanner.Text())
	}
	var rocksPart2 [rockLength]string
	copy(rocksPart2[:], rocks)
	part1day14(rocks)
	part2day14(rocksPart2)
}

func part1day14(rocks []string) {
	var doneMoving bool
	for !doneMoving {
		doneMoving = true
		for i := 1; i < len(rocks); i++ {
			var currentRow string
			var previousRow string
			for j := 0; j < len(rocks[i]); j++ {
				if rocks[i-1][j] == '.' && rocks[i][j] == 'O' {
					previousRow += "O"
					currentRow += "."
					doneMoving = false
				} else {
					currentRow += string(rocks[i][j])
					previousRow += string(rocks[i-1][j])
				}
			}
			rocks[i] = currentRow
			rocks[i-1] = previousRow
		}
	}
	var totalLoad int
	for i := 0; i < len(rocks); i++ {
		totalLoad += strings.Count(string(rocks[i]), "O") * (len(rocks) - i)
	}
	fmt.Println(totalLoad)
}

func part2day14(rocks [rockLength]string) {
	var k int
	var differenceBetweenLoops int
	previousStates := make(map[[rockLength]string]int)
	previousStates[rocks] = 0
	var loopFound bool
	for k < 1000000000*4 {
		direction := dirMap[k%4]
		switch direction {
		case "N":
			{
				var doneMoving bool
				for !doneMoving {
					doneMoving = true
					for i := 1; i < len(rocks); i++ {
						var currentRow string
						var previousRow string
						for j := 0; j < len(rocks[i]); j++ {
							if rocks[i-1][j] == '.' && rocks[i][j] == 'O' {
								previousRow += "O"
								currentRow += "."
								doneMoving = false
							} else {
								currentRow += string(rocks[i][j])
								previousRow += string(rocks[i-1][j])
							}
						}
						rocks[i] = currentRow
						rocks[i-1] = previousRow
					}
				}
			}
		case "E":
			{
				var doneMoving bool
				for !doneMoving {
					doneMoving = true
					for i := 0; i < len(rocks); i++ {
						var currentRow string
						var movedLast bool
						for j := len(rocks[i]) - 1; j > 0; j-- {
							if rocks[i][j-1] == 'O' && rocks[i][j] == '.' {
								currentRow = "O" + currentRow
								movedLast = true
								doneMoving = false
							} else if !movedLast {
								currentRow = string(rocks[i][j]) + currentRow
							} else if rocks[i][j-1] == 'O' {
								currentRow = "O" + currentRow
							} else {
								currentRow = "." + currentRow
								movedLast = false
							}
						}
						if !movedLast {
							rocks[i] = string(rocks[i][0]) + currentRow
						} else {
							rocks[i] = "." + currentRow
						}
					}
				}
			}
		case "S":
			{
				var doneMoving bool
				for !doneMoving {
					doneMoving = true
					for i := len(rocks) - 1; i > 0; i-- {
						var currentRow string
						var previousRow string
						for j := 0; j < len(rocks[i]); j++ {
							if rocks[i-1][j] == 'O' && rocks[i][j] == '.' {
								previousRow += "."
								currentRow += "O"
								doneMoving = false
							} else {
								currentRow += string(rocks[i][j])
								previousRow += string(rocks[i-1][j])
							}
						}
						rocks[i] = currentRow
						rocks[i-1] = previousRow
					}
				}
			}
		case "W":
			{
				var doneMoving bool
				for !doneMoving {
					doneMoving = true
					for i := 0; i < len(rocks); i++ {
						var currentRow string
						var movedLast bool
						for j := 0; j < len(rocks[i])-1; j++ {
							if rocks[i][j+1] == 'O' && rocks[i][j] == '.' {
								currentRow += "O"
								movedLast = true
								doneMoving = false
							} else if !movedLast {
								currentRow += string(rocks[i][j])
							} else if rocks[i][j+1] == 'O' {
								currentRow += "O"
							} else {
								currentRow += "."
								movedLast = false
							}
						}
						if !movedLast {
							rocks[i] = currentRow + string(rocks[i][len(rocks[i])-1])
						} else {
							rocks[i] = currentRow + "."
						}
					}
				}
			}
		}
		k++
		if direction == "E" && !loopFound {
			if val, ok := previousStates[rocks]; ok {
				differenceBetweenLoops = k/4 - val
				loopFound = true
				k = (1000000000 - (1000000000-k/4)%differenceBetweenLoops) * 4
				fmt.Println(k)
			} else {
				previousStates[rocks] = k / 4
			}
		}
	}
	fmt.Println(differenceBetweenLoops)
	var totalLoad int
	for i := 0; i < len(rocks); i++ {
		totalLoad += strings.Count(string(rocks[i]), "O") * (len(rocks) - i)
	}
	fmt.Println(totalLoad)
}
