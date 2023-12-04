package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func Day4() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day4Sample.txt"
	} else {
		fileName = "inputfiles/Day4.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var cards []string
	cardMap := make(map[int]int)
	var cardNumber int
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ":")
		cards = append(cards, strings.Trim(card[1], " "))
		cardMap[cardNumber] = 1
		cardNumber++
	}
	var totalPoints int
	for cardNumber, val := range cards {
		var cardPoints int
		splitCard := strings.Split(val, "|")
		winningNumbers := strings.Split(splitCard[0], " ")
		currentNumbers := strings.Split(splitCard[1], " ")
		numberMap := make(map[string]interface{})
		var numberOfWins int
		for _, val := range currentNumbers {
			numberMap[val] = ""
		}
		for _, winner := range winningNumbers {
			if winner != "" {
				if _, ok := numberMap[winner]; ok {
					cardPoints = int(math.Max(1, float64(cardPoints*2)))
					numberOfWins++
					cardMap[cardNumber+numberOfWins] += cardMap[cardNumber]
				}
			}
		}
		totalPoints += cardPoints
	}
	var totalScratchCards int
	for _, val := range cardMap {
		totalScratchCards += val
	}
	fmt.Println(totalPoints)
	fmt.Println(totalScratchCards)
}
