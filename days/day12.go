package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day12() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day12Sample.txt"
	} else {
		fileName = "inputfiles/Day12.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var springSymbols []string
	var springGroups [][]int
	for scanner.Scan() {
		springs := strings.Split(scanner.Text(), " ")
		springSymbols = append(springSymbols, springs[0])
		var springGrouping []int
		for _, springGroup := range strings.Split(springs[1], ",") {
			springInt, _ := strconv.Atoi(springGroup)
			springGrouping = append(springGrouping, springInt)
		}
		springGroups = append(springGroups, springGrouping)
	}
	var skippedRows int
	var possibleCombinations int
	for i, springs := range springSymbols {
		var numberBroken int
		minimumLength := len(springGroups[i]) - 1
		for _, group := range springGroups[i] {
			numberBroken += group
		}
		trimmedSprings := strings.Trim(springs, ".")
		if len(trimmedSprings) == minimumLength+numberBroken {
			possibleCombinations++
		} else if strings.Count(trimmedSprings, "#") == numberBroken {
			possibleCombinations++
		} else if strings.Count(trimmedSprings, "#")+strings.Count(trimmedSprings, "?") == numberBroken {
			possibleCombinations++
		} else {
			fmt.Println(springs)
			fmt.Println(springGroups[i])
			skippedRows++
		}
	}
	fmt.Println(possibleCombinations)
	fmt.Println(skippedRows)
}
