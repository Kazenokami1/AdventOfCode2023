package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day6Sample.txt"
	} else {
		fileName = "inputfiles/Day6.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var timeArray []string
	var distanceArray []string
	for scanner.Scan() {
		timeArray = strings.Fields(scanner.Text())
		scanner.Scan()
		distanceArray = strings.Fields(scanner.Text())
	}
	day6part1(timeArray, distanceArray)
	day6part2(timeArray, distanceArray)
}

func day6part1(timeArray []string, distanceArray []string) {
	marginOfError := 1
	for i := 1; i < len(timeArray); i++ {
		var waysToWin int
		time, _ := strconv.Atoi(timeArray[i])
		distance, _ := strconv.Atoi(distanceArray[i])
		for j := 0; j < time; j++ {
			if j*(time-j) > distance {
				waysToWin++
			}
		}
		marginOfError *= waysToWin
	}
	fmt.Println(marginOfError)
}

func day6part2(timeArray []string, distanceArray []string) {
	var timeString string
	var distanceString string
	for i := 1; i < len(timeArray); i++ {
		timeString += timeArray[i]
		distanceString += distanceArray[i]
	}
	time, _ := strconv.Atoi(timeString)
	distance, _ := strconv.Atoi(distanceString)
	for i := 0; i < time; i++ {
		if (time-i)*i > distance {
			//First win gets you to the last win on the opposite side
			//First win = i, Last Win = time - i, # of Wins = Last - First + 1 to include the Last Win
			fmt.Println(time - i - i + 1)
			break
		}
	}
}
