package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day8() {
	day8part1()
	day8part2()
}

func day8part1() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day8Sample.txt"
	} else {
		fileName = "inputfiles/Day8.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	nodeMap := make(map[string][]string)
	for scanner.Scan() {
		node := strings.Split(scanner.Text(), " = ")
		branches := strings.Split(node[1], ", ")
		nodeMap[node[0]] = []string{branches[0][1:], branches[1][0 : len(branches[1])-1]}
	}
	currentNode := "AAA"
	var instructionCount int
	for {
		instruction := instructions[instructionCount%len(instructions)]
		if instruction == 'R' {
			currentNode = nodeMap[currentNode][1]
		} else {
			currentNode = nodeMap[currentNode][0]
		}
		if currentNode == "ZZZ" {
			instructionCount++
			break
		}
		instructionCount++
	}
	fmt.Println(instructionCount)
}

func day8part2() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day8Part2Sample.txt"
	} else {
		fileName = "inputfiles/Day8.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	nodeMap := make(map[string][]string)
	for scanner.Scan() {
		node := strings.Split(scanner.Text(), " = ")
		branches := strings.Split(node[1], ", ")
		nodeMap[node[0]] = []string{branches[0][1:], branches[1][0 : len(branches[1])-1]}
	}
	startingNodes := make(map[int]string)
	var numberOfStartingNodes int
	for k := range nodeMap {
		if k[len(k)-1] == 'A' {
			startingNodes[numberOfStartingNodes] = k
			numberOfStartingNodes++
		}
	}
	repeatingMap := make(map[[2]string]int)
	var repeatingArray []int
	for _, currentNode := range startingNodes {
		startNodeMap := make(map[string][]string)
		var instructionCount int
		var repeating bool
		for !repeating {
			if instructions[instructionCount%len(instructions)] == 'R' {
				currentNode = nodeMap[currentNode][1]
			} else {
				currentNode = nodeMap[currentNode][0]
			}
			if currentNode[len(currentNode)-1] == 'Z' {
				currentInstructions := instructions[instructionCount%len(instructions):] + instructions[0:instructionCount%len(instructions)]
				if list, ok := startNodeMap[currentNode]; ok {
					for _, val := range list {
						if val == currentInstructions {
							repeating = true
							repeatingArray = append(repeatingArray, repeatingMap[[2]string{currentNode, currentInstructions}])
							break
						}
					}
				} else {
					startNodeMap[currentNode] = append(startNodeMap[currentNode], currentInstructions)
					repeatingMap[[2]string{currentNode, currentInstructions}] = instructionCount + 1
				}
			}
			instructionCount++
		}
	}
	fmt.Println(lcm(repeatingArray))
}

func lcm(repeatingArray []int) int {
	leastCommonMultiple := repeatingArray[0]
	for i := 1; i < len(repeatingArray); i++ {
		var lcmFound bool
		lcm := leastCommonMultiple
		for !lcmFound {
			if lcm%leastCommonMultiple == 0 && lcm%repeatingArray[i] == 0 {
				lcmFound = true
				leastCommonMultiple = lcm
			} else {
				lcm += leastCommonMultiple
			}
		}
	}
	return leastCommonMultiple
}
