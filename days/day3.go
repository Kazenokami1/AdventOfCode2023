package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func Day3() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day3Sample.txt"
	} else {
		fileName = "inputfiles/Day3.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var schematicRows []string
	symbolMap := make(map[[2]int]rune)
	var schematicRow int
	for scanner.Scan() {
		schematicRows = append(schematicRows, scanner.Text())
		for col, r := range scanner.Text() {
			if !unicode.IsDigit(r) && r != '.' {
				symbolMap[[2]int{schematicRow, col}] = r
			}
		}
		schematicRow++
	}
	day3part1(schematicRows, symbolMap)
	day3part2(schematicRows, symbolMap)
}

func day3part1(schematicRows []string, symbolMap map[[2]int]rune) {
	var enginePartNumbers []int
	var confirmedPart bool
	var possiblePart string
	for row, sch := range schematicRows {
		if confirmedPart {
			partNumber, err := strconv.Atoi(possiblePart)
			Check(err)
			enginePartNumbers = append(enginePartNumbers, partNumber)
			possiblePart = ""
			confirmedPart = false
		}
		for col, r := range sch {
			if unicode.IsDigit(r) {
				possiblePart += string(r)
				if !confirmedPart {
					if _, ok := symbolMap[[2]int{(int(math.Max(0, float64(row-1)))), (int(math.Max(0, float64(col-1))))}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{int(math.Max(0, float64(row-1))), col}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{int(math.Max(0, float64(row-1))), int(math.Min(float64(len(sch)-1), float64(col+1)))}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{row, int(math.Max(0, float64(col-1)))}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{row, int(math.Min(float64(len(sch)-1), float64(col+1)))}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), int(math.Max(0, float64(col-1)))}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), col}]; ok {
						confirmedPart = true
					} else if _, ok := symbolMap[[2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), int(math.Min(float64(len(sch)-1), float64(col+1)))}]; ok {
						confirmedPart = true
					}
				}
			} else if confirmedPart {
				partNumber, err := strconv.Atoi(possiblePart)
				Check(err)
				enginePartNumbers = append(enginePartNumbers, partNumber)
				possiblePart = ""
				confirmedPart = false
			} else if possiblePart != "" {
				possiblePart = ""
			}
		}
	}
	var totalSum int
	for _, val := range enginePartNumbers {
		totalSum += val
	}
	fmt.Println(totalSum)
}

func day3part2(schematicRows []string, symbolMap map[[2]int]rune) {
	gearNumbers := make(map[[2]int][]int)
	var possibleGear bool
	var gearNumber string
	var asteriskPosition [2]int
	for row, sch := range schematicRows {
		if possibleGear {
			gearInt, err := strconv.Atoi(gearNumber)
			Check(err)
			gearNumbers[asteriskPosition] = append(gearNumbers[asteriskPosition], gearInt)
			gearNumber = ""
			possibleGear = false
		}
		for col, r := range sch {
			if unicode.IsDigit(r) {
				gearNumber += string(r)
				if !possibleGear {
					if symbolMap[[2]int{(int(math.Max(0, float64(row-1)))), (int(math.Max(0, float64(col-1))))}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{(int(math.Max(0, float64(row-1)))), (int(math.Max(0, float64(col-1))))}
					} else if symbolMap[[2]int{int(math.Max(0, float64(row-1))), col}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{int(math.Max(0, float64(row-1))), col}
					} else if symbolMap[[2]int{int(math.Max(0, float64(row-1))), int(math.Min(float64(len(sch)-1), float64(col+1)))}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{int(math.Max(0, float64(row-1))), int(math.Min(float64(len(sch)-1), float64(col+1)))}
					} else if symbolMap[[2]int{row, int(math.Max(0, float64(col-1)))}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{row, int(math.Max(0, float64(col-1)))}
					} else if symbolMap[[2]int{row, int(math.Min(float64(len(sch)-1), float64(col+1)))}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{row, int(math.Min(float64(len(sch)-1), float64(col+1)))}
					} else if symbolMap[[2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), int(math.Max(0, float64(col-1)))}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), int(math.Max(0, float64(col-1)))}
					} else if symbolMap[[2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), col}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), col}
					} else if symbolMap[[2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), int(math.Min(float64(len(sch)-1), float64(col+1)))}] == '*' {
						possibleGear = true
						asteriskPosition = [2]int{int(math.Min(float64(len(schematicRows)-1), float64(row+1))), int(math.Min(float64(len(sch)-1), float64(col+1)))}
					}
				}
			} else if possibleGear {
				gearInt, err := strconv.Atoi(gearNumber)
				Check(err)
				gearNumbers[asteriskPosition] = append(gearNumbers[asteriskPosition], gearInt)
				gearNumber = ""
				possibleGear = false
			} else if gearNumber != "" {
				gearNumber = ""
			}
		}
	}
	var totalSum int
	for _, val := range gearNumbers {
		if len(val) == 2 {
			totalSum += val[0] * val[1]
		}
	}
	fmt.Println(totalSum)
}
