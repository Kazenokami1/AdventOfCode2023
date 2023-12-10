package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var validPipesByDirection = make(map[string]string)

func init() {
	validPipesByDirection["UP"] = "|7FS"
	validPipesByDirection["DOWN"] = "|LJS"
	validPipesByDirection["LEFT"] = "-LFS"
	validPipesByDirection["RIGHT"] = "-7JS"
}

func Day10() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day10Sample.txt"
	} else {
		fileName = "inputfiles/Day10.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var pipeMap []string
	for scanner.Scan() {
		pipeMap = append(pipeMap, scanner.Text())
	}
	pipeMapping := make(map[[2]int]*pipe)
	var startingPipe *pipe
	for i, pipeRow := range pipeMap {
		for j, pipeShape := range pipeRow {
			if pipeMap[i][j] != '.' {
				pipeMapping[[2]int{i, j}] = &pipe{shape: string(pipeShape), position: [2]int{i, j}}
			}
			if pipeShape == 'S' {
				startingPipe = pipeMapping[[2]int{i, j}]
			}
		}
	}
	validPipes := make(map[[2]int]*pipe)
	for k, pipe := range pipeMapping {
		switch pipe.shape {
		case "|":
			if neighbor, ok := pipeMapping[[2]int{k[0] - 1, k[1]}]; ok && strings.Contains(validPipesByDirection["UP"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
			if neighbor, ok := pipeMapping[[2]int{k[0] + 1, k[1]}]; ok && strings.Contains(validPipesByDirection["DOWN"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
		case "L":
			if neighbor, ok := pipeMapping[[2]int{k[0] - 1, k[1]}]; ok && strings.Contains(validPipesByDirection["UP"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
			if neighbor, ok := pipeMapping[[2]int{k[0], k[1] + 1}]; ok && strings.Contains(validPipesByDirection["RIGHT"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
		case "F":
			if neighbor, ok := pipeMapping[[2]int{k[0], k[1] + 1}]; ok && strings.Contains(validPipesByDirection["RIGHT"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
			if neighbor, ok := pipeMapping[[2]int{k[0] + 1, k[1]}]; ok && strings.Contains(validPipesByDirection["DOWN"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
		case "-":
			if neighbor, ok := pipeMapping[[2]int{k[0], k[1] - 1}]; ok && strings.Contains(validPipesByDirection["LEFT"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
			if neighbor, ok := pipeMapping[[2]int{k[0], k[1] + 1}]; ok && strings.Contains(validPipesByDirection["RIGHT"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
		case "7":
			if neighbor, ok := pipeMapping[[2]int{k[0], k[1] - 1}]; ok && strings.Contains(validPipesByDirection["LEFT"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
			if neighbor, ok := pipeMapping[[2]int{k[0] + 1, k[1]}]; ok && strings.Contains(validPipesByDirection["DOWN"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
		case "J":
			if neighbor, ok := pipeMapping[[2]int{k[0] - 1, k[1]}]; ok && strings.Contains(validPipesByDirection["UP"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
			if neighbor, ok := pipeMapping[[2]int{k[0], k[1] - 1}]; ok && strings.Contains(validPipesByDirection["LEFT"], neighbor.shape) {
				pipe.addNeighbor(neighbor)
			}
		case "S":
			startingPipe = pipe
		}
		if len(pipe.neighbors) == 2 {
			validPipes[pipe.position] = pipe
		}
	}
	var notChanged bool
	for !notChanged {
		notChanged = true
		for k, pipe := range validPipes {
			for _, neighbor := range pipe.neighbors {
				if _, ok := validPipes[neighbor.position]; !ok && neighbor != startingPipe {
					delete(validPipes, k)
					notChanged = false
				}
			}
		}
	}
	var possibleS string
	if pipe, ok := validPipes[[2]int{startingPipe.position[0] + 1, startingPipe.position[1]}]; ok && strings.Contains(validPipesByDirection["DOWN"], pipe.shape) {
		startingPipe.neighbors = append(startingPipe.neighbors, pipe)
		possibleS += "|7F"
	}
	if pipe, ok := validPipes[[2]int{startingPipe.position[0] - 1, startingPipe.position[1]}]; ok && strings.Contains(validPipesByDirection["UP"], pipe.shape) {
		startingPipe.neighbors = append(startingPipe.neighbors, pipe)
		possibleS += "|LJ"
	}
	if pipe, ok := validPipes[[2]int{startingPipe.position[0], startingPipe.position[1] + 1}]; ok && strings.Contains(validPipesByDirection["RIGHT"], pipe.shape) {
		startingPipe.neighbors = append(startingPipe.neighbors, pipe)
		possibleS += "LF-"
	}
	if pipe, ok := validPipes[[2]int{startingPipe.position[0], startingPipe.position[1] - 1}]; ok && strings.Contains(validPipesByDirection["LEFT"], pipe.shape) {
		startingPipe.neighbors = append(startingPipe.neighbors, pipe)
		possibleS += "J7-"
	}
	for i, sValue := range possibleS {
		if strings.Contains(possibleS[i+1:], string(sValue)) {
			startingPipe.shape = string(sValue)
		}
	}
	validPipes[startingPipe.position] = startingPipe
	var loopComplete bool
	currentPipe := startingPipe
	previousPipe := startingPipe.neighbors[0]
	mainLoopPipes := make(map[[2]int]*pipe)
	for !loopComplete {
		for _, nextPipe := range currentPipe.neighbors {
			if nextPipe != previousPipe {
				previousPipe = currentPipe
				currentPipe = nextPipe
				mainLoopPipes[currentPipe.position] = currentPipe
				break
			}
		}
		if currentPipe == startingPipe {
			loopComplete = true
		}
	}
	fmt.Println(len(mainLoopPipes) / 2)
	var totalEnclosed int
	for i := 0; i < len(pipeMap); i++ {
		var pipeEnclosed bool
		var loopEdge string
		for j := 0; j < len(pipeMap[i]); j++ {
			position := [2]int{i, j}
			if pipe, ok := mainLoopPipes[position]; ok {
				if pipe.shape == "|" {
					pipeEnclosed = !pipeEnclosed
					loopEdge = ""
				} else if loopEdge == "" {
					loopEdge = pipe.shape
				} else {
					if (loopEdge == "L" && pipe.shape == "7") || (loopEdge == "F" && pipe.shape == "J") {
						pipeEnclosed = !pipeEnclosed
						loopEdge = ""
					} else if pipe.shape != "-" {
						loopEdge = ""
					}
				}
				fmt.Print(pipe.shape)
			} else if pipeEnclosed {
				totalEnclosed++
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println(totalEnclosed)
}
