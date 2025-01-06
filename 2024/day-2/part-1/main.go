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

func checkLevelTendency(level []int) bool {
	// true - rising tendency, false - decreasing tendency
	var tendency bool

	// define tendency based on first two elements
	if level[0] > level[1] {
		tendency = true
	} else {
		tendency = false
	}

	for i := 1; i < len(level); i++ {
		// catch if next elements are equal
		// if so there is no tendency
		if level[i-1] == level[i] {
			return false
		}

		if level[i-1] > level[i] {
			if !tendency {
				// if tendency is not rising but level[i-1] is bigger than level[i]
				// therefore there is no tendency and whole raport is invalid
				return false
			}
		}
		if level[i-1] < level[i] {
			if tendency {
				// if tendency is rising but level[i-1] is smaller than level[i]
				// therefore there is no tendency and whole raport is invalid
				return false
			}
		}
	}

	return true
}

func checkLevelDifference(level []int) bool {
	for i := 1; i < len(level); i++ {
		diff := calculateDifference(level[i-1], level[i])
		if diff == 0 {
			return false
		}
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func main() {
	/*
		1. Input data is a slice of slices [][]int. Each row is a report which is a slice []int. Every report contains levels, each level is one element in slice
		2. Check if levels are either all increasing or all decreasing
			2a. If not omit that report
		3. Check adjacent level difference, maximum acceptable diff is from 1 to 3
			3a. If difference is less then 1 (for example the number stays the same) or greated than 4, omit that report
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
		}
	}
	fmt.Println(safeReports)
}
