package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day15() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day15Sample.txt"
	} else {
		fileName = "inputfiles/Day15.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	var initialization string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		initialization += scanner.Text()
	}
	initArray := strings.Split(initialization, ",")
	part1Day15(initArray)
	part2Day15(initArray)
}

func part1Day15(initArray []string) {
	var initValue int
	for _, init := range initArray {
		var valueToAdd int
		for _, char := range init {
			valueToAdd = ((valueToAdd + int(char)) * 17) % 256
		}
		initValue += valueToAdd
	}
	fmt.Println(initValue)
}

func part2Day15(initArray []string) {
	boxMap := make(map[int]*box)
	for _, init := range initArray {
		var boxNumber int
		for i, char := range init {
			switch char {
			case '-':
				if existingBox, ok := boxMap[boxNumber]; ok {
					lens := init[0:i]
					for j, slot := range existingBox.slots {
						if strings.Contains(slot, lens) {
							if j < len(boxMap[boxNumber].slots)-1 {
								existingBox.slots = append(existingBox.slots[0:j], existingBox.slots[j+1:]...)
							} else {
								existingBox.slots = existingBox.slots[0:j]
							}
						}
					}
				}
			case '=':
				if existingBox, ok := boxMap[boxNumber]; ok {
					var lensFound bool
					for j, slot := range existingBox.slots {
						if strings.Contains(slot, init[0:i]) {
							existingBox.slots[j] = init[0:i] + " " + init[i+1:]
							lensFound = true
						}
					}
					if !lensFound {
						existingBox.slots = append(existingBox.slots, init[0:i]+" "+init[i+1:])
					}
				} else {
					boxMap[boxNumber] = &box{slots: []string{init[0:i] + " " + init[i+1:]}}
				}
			default:
				boxNumber = ((boxNumber + int(char)) * 17) % 256
			}
		}
	}
	var totalFocusingPower int
	for boxNumber, mappedBox := range boxMap {
		for lensNumber, slot := range mappedBox.slots {
			focalLength, _ := strconv.Atoi(slot[len(slot)-1:])
			totalFocusingPower += (boxNumber + 1) * (lensNumber + 1) * focalLength
		}
	}
	fmt.Println(totalFocusingPower)
}
