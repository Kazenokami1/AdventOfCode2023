package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day9Sample.txt"
	} else {
		fileName = "inputfiles/Day9.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var historyValues [][]int
	for scanner.Scan() {
		var history []int
		value := strings.Split(scanner.Text(), " ")
		for _, val := range value {
			intVal, _ := strconv.Atoi(val)
			history = append(history, intVal)
		}
		historyValues = append(historyValues, history)
	}
	var totalDiff int
	var totalPrevDiff int
	for _, historyValue := range historyValues {
		var allZeroes bool
		var historyDifferences [][]int
		historyDifferences = append(historyDifferences, historyValue)
		for !allZeroes {
			allZeroes = true
			var differences []int
			for i := 0; i < len(historyDifferences[len(historyDifferences)-1])-1; i++ {
				difference := historyDifferences[len(historyDifferences)-1][i+1] - historyDifferences[len(historyDifferences)-1][i]
				if difference != 0 {
					allZeroes = false
				}
				differences = append(differences, difference)
			}
			historyDifferences = append(historyDifferences, differences)
		}
		for i := len(historyDifferences) - 2; i >= 0; i-- {
			prediction := historyDifferences[i][len(historyDifferences[i])-1] + historyDifferences[i+1][len(historyDifferences[i+1])-1]
			historyDifferences[i] = append(historyDifferences[i], prediction)
			if i == 0 {
				totalDiff += prediction
			}
		}
		for i := len(historyDifferences) - 2; i >= 0; i-- {
			prediction := historyDifferences[i][0] - historyDifferences[i+1][0]
			historyDifferences[i] = append([]int{prediction}, historyDifferences[i]...)
			if i == 0 {
				totalPrevDiff += prediction
			}
		}
	}
	fmt.Println(totalDiff)
	fmt.Println(totalPrevDiff)
}
