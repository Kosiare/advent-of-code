package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseInput(path string) (data []string, err error) {
	readFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		data = append(data, fileScanner.Text())
	}

	return data, nil
}

func main() {
	var clearInput []string
	var sum int

	filePath := ""
	flag.StringVar(&filePath, "i", "", "input file to parse")
	flag.Parse()

	input, err := parseInput(filePath)
	if err != nil {
		log.Fatalf("FATAL: error reading file:\n%v", err)
	}

	expr := "(mul\\([0-9]{1,3},[0-9]{1,3}\\))|do\\(\\)|don't\\(\\)"
	r, _ := regexp.Compile(expr)

	expr2 := "[0-9]{1,3}"
	r2, _ := regexp.Compile(expr2)

	for i := range input {
		clearInput = append(clearInput, r.FindAllString(input[i], -1)...)
	}

	active := true
	for i := 0; i < len(clearInput); i++ {
		if clearInput[i] == "don't()" {
			active = false
			continue
		}
		if clearInput[i] == "do()" {
			active = true
			continue
		}

		if active {
			nums := r2.FindAllString(clearInput[i], -1)
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			sum += (num1 * num2)
		}

	}

	fmt.Println(sum)

}
