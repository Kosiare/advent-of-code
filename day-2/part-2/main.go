package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) (input [][]int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		input = append(input, make([]int, len(parts)))
		for p := range parts {
			value, e := strconv.Atoi(parts[p])
			if e != nil {
				return [][]int{}, e
			}
			input[lineNumber][p] = value
		}
		lineNumber++
	}
	return
}

func calculateDifference(left int, right int) (difference int) {
	difference = left - right
	if difference < 0 {
		difference = difference * -1
	}
	return difference
}

func checkLevelTendency(report []int) bool {
	// true - rising tendency, false - decreasing tendency
	var tendency bool

	// define tendency based on first two elements
	if report[0] > report[1] {
		tendency = true
	} else {
		tendency = false
	}

	for i := 1; i < len(report); i++ {
		// catch if next elements are equal
		// if so there is no tendency
		if report[i-1] == report[i] {
			return false
		}

		if report[i-1] > report[i] {
			if !tendency {
				// if tendency is not rising but level[i-1] is bigger than level[i]
				// therefore there is no tendency and whole raport is invalid
				return false
			}
		}
		if report[i-1] < report[i] {
			if tendency {
				// if tendency is rising but level[i-1] is smaller than level[i]
				// therefore there is no tendency and whole raport is invalid
				return false
			}
		}
	}

	return true
}

func checkLevelDifference(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := calculateDifference(report[i-1], report[i])
		if diff == 0 {
			return false
		}
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func deleteFaultyLevel(report []int, indexToDelete int) (trimmedReport []int) {
	for i := 0; i < len(report); i++ {
		if i != indexToDelete {
			trimmedReport = append(trimmedReport, report[i])
		}
	}
	return trimmedReport
}

func main() {
	/*
		1. Input data is a slice of slices [][]int. Each row is a report which is a slice []int. Every report contains levels, each level is one element in slice
		2. Check if levels are either all increasing or all decreasing
			2a. If tendency has been broken try to delete FIRST level and check that report again
			2b. If tendency is still broken try to delete first FAULTY level and check that report again - check it for whole pair:
				1) Delete left level from pair and check it
				2) Delete right level from pair and check it
			2c. If it is still faulty omit that report completely
		3. Check adjacent level difference, maximum acceptable diff is from 1 to 3
			3a. If difference is in not acceptable range try to delete FIRST level and check that report again
			3b. If it still faulty try to delete first FAULTY level and check that report again
				1) Delete left level from pair and check it
				2) Delete right level from pair and check it
			3c. If still difference is not in range omit that report completely
		4. If certain level passes previous two conditions then report is considered safe
		5. Check how many safe reports are there
	*/

	var input [][]int

	path := ""
	flag.StringVar(&path, "i", "", "input file to parse")
	flag.Parse()

	input, err := parseInput(path)
	if err != nil {
		log.Fatalln("Cannot parse input:", err)
	}

	safeReports := 0
	for i := range input {
		if checkLevelTendency(input[i]) && checkLevelDifference(input[i]) {
			safeReports++
			continue
		}

		for j := 0; j < len(input[i]); j++ {
			trimmedInput := deleteFaultyLevel(input[i], j)
			if checkLevelTendency(trimmedInput) && checkLevelDifference(trimmedInput) {
				safeReports++
				break
			}
		}

	}
	fmt.Println(safeReports)

}
