package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day2Sample.txt"
	} else {
		fileName = "inputfiles/Day2.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var totalCubePower int
	for scanner.Scan() {
		cubePower := 1
		cubeCounts := strings.FieldsFunc(strings.ReplaceAll(scanner.Text(), ",", ""), day2split)
		var minCubes = make(map[string]int)
		for _, cubes := range cubeCounts {
			if strings.Contains(cubes, "Game") {
				continue
			} else {
				numbers := strings.Split(strings.Trim(cubes, " "), " ")
				for i := 1; i < len(numbers); i += 2 {
					cubeCount, err := strconv.Atoi(numbers[i-1])
					Check(err)
					if val, ok := minCubes[numbers[i]]; !ok {
						minCubes[numbers[i]] = cubeCount
					} else if cubeCount > val {
						minCubes[numbers[i]] = cubeCount
					}
				}
			}
		}
		for _, v := range minCubes {
			cubePower *= v
		}
		totalCubePower += cubePower
	}
	fmt.Println(totalCubePower)
}

func day2split(r rune) bool {
	return r == ':' || r == ';'
}
