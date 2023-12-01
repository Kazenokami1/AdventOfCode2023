package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var intStrings = make(map[string]string)

func init() {
	intStrings["one"] = "1"
	intStrings["two"] = "2"
	intStrings["three"] = "3"
	intStrings["four"] = "4"
	intStrings["five"] = "5"
	intStrings["six"] = "6"
	intStrings["seven"] = "7"
	intStrings["eight"] = "8"
	intStrings["nine"] = "9"
	intStrings["zero"] = "0"
}

func Day1() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day1Sample.txt"
	} else {
		fileName = "inputfiles/Day1.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var totalCalibrationValue int
	for scanner.Scan() {
		var calibrationValue string
		calibrationText := scanner.Text()
		for i := 0; i < len(calibrationText); i++ {
			_, err := strconv.Atoi(string(calibrationText[i]))
			if err != nil {
				for k := range intStrings {
					if strings.Contains(calibrationText[0:i+1], k) {
						calibrationValue = intStrings[k]
						break
					}
				}
				if len(calibrationValue) == 1 {
					break
				} else {
					continue
				}
			} else {
				calibrationValue = string(calibrationText[i])
				break
			}
		}
		for i := len(calibrationText) - 1; i >= 0; i-- {
			_, err := strconv.Atoi(string(calibrationText[i]))
			if err != nil {
				for k := range intStrings {
					if strings.Contains(calibrationText[i:], k) {
						calibrationValue += intStrings[k]
						break
					}
				}
				if len(calibrationValue) == 2 {
					break
				} else {
					continue
				}
			} else {
				calibrationValue += string(calibrationText[i])
				break
			}
		}
		value, err := strconv.Atoi(calibrationValue)
		Check(err)
		totalCalibrationValue += value
	}
	fmt.Println(totalCalibrationValue)
}
