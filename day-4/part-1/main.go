package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func lookUpHorizontal(input []string) int {
	ans := 0
	for i := range input {
		// normal
		expr := "XMAS"
		r, _ := regexp.Compile(expr)
		ans += len(r.FindAllString(input[i], -1))

		// backwards
		expr = "SAMX"
		r2, _ := regexp.Compile(expr)
		ans += len(r2.FindAllString(input[i], -1))
	}
	return ans
}

func lookUpVertical(input []string) int {
	ans := 0
	colLen := len(input)
	for col := range input {
		for char := range input[col] {
			// normal
			if (col+3 < colLen) && (string(input[col][char]) == "X") && (string(input[col+1][char]) == "M") && (string(input[col+2][char]) == "A") && (string(input[col+3][char]) == "S") {
				ans++
			}
			// backwards
			if (col+3 < colLen) && (string(input[col][char]) == "S") && (string(input[col+1][char]) == "A") && (string(input[col+2][char]) == "M") && (string(input[col+3][char]) == "X") {
				ans++
			}
		}
	}
	return ans
}

func lookUpDiagonal(input []string) int {
	ans := 0
	colLen := len(input)
	rowLen := len(input[0])

	for col := range input {
		for char := range input[col] {
			// normal - left to right
			if (col+3 < colLen && char+3 < rowLen) && (string(input[col][char]) == "X") && (string(input[col+1][char+1]) == "M") && (string(input[col+2][char+2]) == "A") && (string(input[col+3][char+3]) == "S") {
				ans++
			}
			// backwards - left to right
			if (col+3 < colLen && char+3 < rowLen) && (string(input[col][char]) == "S") && (string(input[col+1][char+1]) == "A") && (string(input[col+2][char+2]) == "M") && (string(input[col+3][char+3]) == "X") {
				ans++
			}
			// normal - right to left
			if (col-3 >= 0 && char+3 < rowLen) && (string(input[col][char]) == "X") && (string(input[col-1][char+1]) == "M") && (string(input[col-2][char+2]) == "A") && (string(input[col-3][char+3]) == "S") {
				ans++
			}
			// backwards - right to left
			if (col-3 >= 0 && char+3 < rowLen) && (string(input[col][char]) == "S") && (string(input[col-1][char+1]) == "A") && (string(input[col-2][char+2]) == "M") && (string(input[col-3][char+3]) == "X") {
				ans++
			}
		}
	}

	return ans
}

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
	var sum int

	filePath := ""
	flag.StringVar(&filePath, "i", "", "input file to parse")
	flag.Parse()

	input, err := parseInput(filePath)
	if err != nil {
		log.Fatalf("FATAL: error reading file:\n%v", err)
	}

	sum += lookUpHorizontal(input)
	sum += lookUpVertical(input)
	sum += lookUpDiagonal(input)

	fmt.Println(sum)
}
