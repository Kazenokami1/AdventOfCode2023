package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day5Sample.txt"
	} else {
		fileName = "inputfiles/Day5.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var almanac []string
	for scanner.Scan() {
		almanac = append(almanac, scanner.Text())
	}
	for i, instruction := range almanac {
		if instruction == "" {
			almanac = append(almanac[0:i], almanac[i+1:]...)
		}
	}
	seedsString := strings.Split(almanac[0], " ")
	var seeds []int
	for _, val := range seedsString {
		seedNumber, err := strconv.Atoi(val)
		if err != nil {
			continue
		} else {
			seeds = append(seeds, seedNumber)
		}
	}
	almanac = almanac[1:]
	almanacMaps := make(map[int][][]int)
	mapNumber := -1
	for _, val := range almanac {
		if strings.Contains(val, ":") {
			mapNumber++
		} else {
			var intArray []int
			splitVal := strings.Split(val, " ")
			for _, number := range splitVal {
				intVal, _ := strconv.Atoi(number)
				intArray = append(intArray, intVal)
			}
			almanacMaps[mapNumber] = append(almanacMaps[mapNumber], intArray)
		}
	}
	day5part1(seeds, almanacMaps)
	day5part2(seeds, almanacMaps)
}

func day5part1(seeds []int, almanacMaps map[int][][]int) {
	var nearestLocation int
	for _, val := range seeds {
		for i := 0; i < len(almanacMaps); i++ {
			var mapped bool
			for j := 0; j < len(almanacMaps[i]); j++ {
				if val >= almanacMaps[i][j][1] && val < almanacMaps[i][j][1]+almanacMaps[i][j][2] && !mapped {
					val = val - (almanacMaps[i][j][1] - almanacMaps[i][j][0])
					mapped = true
					j = len(almanacMaps[i])
				}
			}
		}
		if val < nearestLocation || nearestLocation == 0 {
			nearestLocation = val
		}
	}
	fmt.Println(nearestLocation)
}

func day5part2(seeds []int, almanacMaps map[int][][]int) {
	var seedRanges [][]int
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, seeds[i:i+2])
		// Highest Seed Inclusive
		seedRanges[len(seedRanges)-1][1] += seedRanges[len(seedRanges)-1][0] - 1
	}
	var nearestLocation int
	for i := 0; i < len(almanacMaps); i++ {
		seedRanges = convertSeeds(seedRanges, almanacMaps[i])
	}
	for i := 0; i < len(seedRanges); i++ {
		if seedRanges[i][0] < nearestLocation || i == 0 {
			nearestLocation = seedRanges[i][0]
		}
	}
	fmt.Println(nearestLocation)
}

func convertSeeds(seedRanges [][]int, almanacMap [][]int) [][]int {
	for i := 0; i < len(seedRanges); i++ {
		var mapped bool
		for _, convert := range almanacMap {
			if seedRanges[i][0] < convert[1] && seedRanges[i][1] >= convert[1] && !mapped {
				seedRanges = append(seedRanges, []int{seedRanges[i][0], convert[1] - 1})
				seedRanges[i][0] = convert[1]
			}
			if seedRanges[i][0] < convert[1]+convert[2]-1 && seedRanges[i][1] > convert[1]+convert[2]-1 && !mapped {
				seedRanges = append(seedRanges, []int{convert[1] + convert[2], seedRanges[i][1]})
				seedRanges[i][1] = convert[1] + convert[2] - 1
			}
			if seedRanges[i][0] >= convert[1] && seedRanges[i][0] < convert[1]+convert[2] && !mapped {
				seedRanges[i][0] -= (convert[1] - convert[0])
				seedRanges[i][1] -= (convert[1] - convert[0])
				mapped = true
			}
		}
	}
	return seedRanges
}
