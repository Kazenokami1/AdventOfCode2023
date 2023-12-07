package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cardRankingsPart1 = "23456789TJQKA"
var cardRankingsPart2 = "J23456789TQKA"

func Day7() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day7Sample.txt"
	} else {
		fileName = "inputfiles/Day7.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	cardHands := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cards := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(cards[1])
		cardHands[cards[0]] = bid
	}
	day7part1(cardHands)
	day7part2(cardHands)
}

func day7part1(cardHands map[string]int) {
	var totalScore int
	var orderedHands []string
	for hand := range cardHands {
		if len(orderedHands) == 0 {
			orderedHands = append(orderedHands, hand)
		} else {
			var losingHand bool
			var currentHand int
			for !losingHand && currentHand < len(orderedHands) {
				losingHand = compareHandsPart1(hand, orderedHands[currentHand])
				if !losingHand {
					currentHand++
				}
			}
			var newOrderedHands []string
			for i := 0; i < currentHand; i++ {
				newOrderedHands = append(newOrderedHands, orderedHands[i])
			}
			newOrderedHands = append(newOrderedHands, hand)
			orderedHands = append(newOrderedHands, orderedHands[currentHand:]...)
		}
	}
	for i := 0; i < len(orderedHands); i++ {
		totalScore += cardHands[orderedHands[i]] * (i + 1)
	}
	fmt.Println(totalScore)
}

func day7part2(cardHands map[string]int) {
	var totalScore int
	var orderedHands []string
	for hand := range cardHands {
		if len(orderedHands) == 0 {
			orderedHands = append(orderedHands, hand)
		} else {
			var losingHand bool
			var currentHand int
			for !losingHand && currentHand < len(orderedHands) {
				losingHand = compareHandsPart2(hand, orderedHands[currentHand])
				if !losingHand {
					currentHand++
				}
			}
			var newOrderedHands []string
			for i := 0; i < currentHand; i++ {
				newOrderedHands = append(newOrderedHands, orderedHands[i])
			}
			newOrderedHands = append(newOrderedHands, hand)
			orderedHands = append(newOrderedHands, orderedHands[currentHand:]...)
		}
	}
	for i := 0; i < len(orderedHands); i++ {
		totalScore += cardHands[orderedHands[i]] * (i + 1)
	}
	fmt.Println(totalScore)
}

func compareHandsPart1(handOne string, handTwo string) bool {
	handOneMap := make(map[rune]int)
	handTwoMap := make(map[rune]int)
	for _, val := range handOne {
		handOneMap[val]++
	}
	for _, val := range handTwo {
		handTwoMap[val]++
	}
	if len(handOneMap) < len(handTwoMap) {
		return false
	} else if len(handOneMap) > len(handTwoMap) {
		return true
	} else {
		var maxHandOne int
		var maxHandTwo int
		for _, matches := range handOneMap {
			if matches > maxHandOne {
				maxHandOne = matches
			}
		}
		for _, matches := range handTwoMap {
			if matches > maxHandTwo {
				maxHandTwo = matches
			}
		}
		if maxHandOne > maxHandTwo {
			return false
		} else if maxHandOne < maxHandTwo {
			return true
		} else {
			for i := 0; i < len(handOne); i++ {
				handOneRank := strings.Index(cardRankingsPart1, string(handOne[i]))
				handTwoRank := strings.Index(cardRankingsPart1, string(handTwo[i]))
				if handOneRank > handTwoRank {
					return false
				} else if handOneRank < handTwoRank {
					return true
				}
			}
		}
	}
	return false
}

func compareHandsPart2(handOne string, handTwo string) bool {
	handOneMap := make(map[rune]int)
	handTwoMap := make(map[rune]int)
	var handOneJokers int
	var handTwoJokers int
	for _, val := range handOne {
		if val != 'J' {
			handOneMap[val]++
		} else {
			handOneJokers++
		}
	}
	for _, val := range handTwo {
		if val != 'J' {
			handTwoMap[val]++
		} else {
			handTwoJokers++
		}
	}
	if len(handOneMap) == 0 {
		handOneMap['J'] = 0
	}
	if len(handTwoMap) == 0 {
		handTwoMap['J'] = 0
	}
	if len(handOneMap) < len(handTwoMap) {
		return false
	} else if len(handOneMap) > len(handTwoMap) {
		return true
	} else {
		var maxHandOne int
		var maxHandTwo int
		for _, matches := range handOneMap {
			if matches > maxHandOne {
				maxHandOne = matches
			}
		}
		maxHandOne += handOneJokers
		for _, matches := range handTwoMap {
			if matches > maxHandTwo {
				maxHandTwo = matches
			}
		}
		maxHandTwo += handTwoJokers
		if maxHandOne > maxHandTwo {
			return false
		} else if maxHandOne < maxHandTwo {
			return true
		} else {
			for i := 0; i < len(handOne); i++ {
				handOneRank := strings.Index(cardRankingsPart2, string(handOne[i]))
				handTwoRank := strings.Index(cardRankingsPart2, string(handTwo[i]))
				if handOneRank > handTwoRank {
					return false
				} else if handOneRank < handTwoRank {
					return true
				}
			}
		}
	}
	return false
}
