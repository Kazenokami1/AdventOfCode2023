package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day19() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day19Sample.txt"
	} else {
		fileName = "inputfiles/Day19.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	workflowMap := make(map[string]string)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		workflow := strings.Split(scanner.Text(), "{")
		workflowMap[workflow[0]] = workflow[1][0 : len(workflow[1])-1]
	}
	var parts []part
	for scanner.Scan() {
		partVals := strings.Split(scanner.Text()[1:len(scanner.Text())-1], ",")
		xVal, _ := strconv.Atoi(partVals[0][2:])
		mVal, _ := strconv.Atoi(partVals[1][2:])
		aVal, _ := strconv.Atoi(partVals[2][2:])
		sVal, _ := strconv.Atoi(partVals[3][2:])
		partToAdd := part{xVal: xVal, mVal: mVal, aVal: aVal, sVal: sVal}
		partToAdd.calcPartValue()
		parts = append(parts, partToAdd)
	}
	var acceptedParts []part
	for _, part := range parts {
		accepted := testWorkFlow(workflowMap["in"], part, workflowMap)
		if accepted {
			acceptedParts = append(acceptedParts, part)
		}
	}
	var totalPartsValues int
	for _, part := range acceptedParts {
		totalPartsValues += part.partValue
	}
	fmt.Println(totalPartsValues)
}

func testWorkFlow(workflow string, testPart part, workflowMap map[string]string) bool {
	var accepted bool
	workflowSteps := strings.Split(workflow, ",")
	for _, test := range workflowSteps {
		var passedTest bool
		var testAgainst int
		colonIndex := strings.Index(test, ":")
		if colonIndex >= 0 {
			testAgainst, _ = strconv.Atoi(test[2:colonIndex])
		} else {
			if workflow, ok := workflowMap[test]; ok {
				return testWorkFlow(workflow, testPart, workflowMap)
			} else {
				return test == "A"
			}
		}
		switch test[0] {
		case 'x':
			if test[1] == '>' {
				passedTest = isGreaterThan(testPart.xVal, testAgainst)
			} else {
				passedTest = isGreaterThan(testAgainst, testPart.xVal)
			}
		case 'm':
			if test[1] == '>' {
				passedTest = isGreaterThan(testPart.mVal, testAgainst)
			} else {
				passedTest = isGreaterThan(testAgainst, testPart.mVal)
			}
		case 'a':
			if test[1] == '>' {
				passedTest = isGreaterThan(testPart.aVal, testAgainst)
			} else {
				passedTest = isGreaterThan(testAgainst, testPart.aVal)
			}
		case 's':
			if test[1] == '>' {
				passedTest = isGreaterThan(testPart.sVal, testAgainst)
			} else {
				passedTest = isGreaterThan(testAgainst, testPart.sVal)
			}
		}
		if passedTest {
			nextWorkflow := test[colonIndex+1:]
			if workflow, ok := workflowMap[nextWorkflow]; ok {
				return testWorkFlow(workflow, testPart, workflowMap)
			} else if nextWorkflow == "A" {
				return true
			} else if nextWorkflow == "R" {
				return false
			}
		}
	}
	return accepted
}
