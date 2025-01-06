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

type arrayFlags []string

func (a *arrayFlags) String() string {
	return strings.Join(*a, " ")
}

func (a *arrayFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func parseInput(path string) (input [][]int, err error) {
	readFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNumber := 0
	for scanner.Scan() {
		var parts []string
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts = strings.Split(line, "|")
		}
		if strings.Contains(line, ",") {
			parts = strings.Split(line, ",")
		}
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

func isFollowingRules(update []int, rules [][]int) bool {
	idx := make(map[int]int)

	for i, num := range update {
		idx[num] = i
	}

	for _, rule := range rules {
		a, b := rule[0], rule[1]
		if aIdx, boolA := idx[a]; boolA {
			if bIdx, boolB := idx[b]; boolB {
				if !(aIdx < bIdx) {
					return false
				}
			}
		}
	}

	return true
}

func sortFaultyUpdate(update []int, rules [][]int) []int {
	for true {
		isSorted := true
		for i := 0; i < len(update)-1; i++ {
			for _, rule := range rules {
				if rule[0] == update[i+1] && rule[1] == update[i] {
					isSorted = false
					update[i], update[i+1] = update[i+1], update[i]
				}
			}
		}
		if isSorted {
			return update
		}
	}

	// this should never execute
	return []int{0, 0, 0}
}

func main() {
	var filePaths arrayFlags
	flag.Var(&filePaths, "i", "input files to parse")
	flag.Parse()

	// parse rule list
	rules, err := parseInput(filePaths[0])
	if err != nil {
		log.Fatalf("FATAL: error reading rules file:\n%v", err)
	}
	// parse updates list
	updates, err := parseInput(filePaths[1])
	if err != nil {
		log.Fatalf("FATAL: error reading updates file:\n%v", err)
	}

	ans := 0
	for i := range updates {
		valid := isFollowingRules(updates[i], rules)
		if !valid {
			sorted := sortFaultyUpdate(updates[i], rules)
			ans += sorted[len(updates[i])/2]
		}

	}

	fmt.Println(ans)
}
