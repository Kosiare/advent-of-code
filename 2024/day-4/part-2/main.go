package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	var sum int

	filePath := ""
	flag.StringVar(&filePath, "i", "", "input file to parse")
	flag.Parse()

	input, err := parseInput(filePath)
	if err != nil {
		log.Fatalf("FATAL: error reading file:\n%v", err)
	}

	colLen := len(input)
	rowLen := len(input[0])
	for col := range input {
		for char := range input[col] {
			// check if there is enough space
			if col+2 < colLen && char+2 < rowLen {
				// M . S
				// . A .
				// M . S
				if (string(input[col][char]) == "M") && (string(input[col+1][char+1]) == "A") && (string(input[col+2][char+2]) == "S") && (string(input[col+2][char]) == "M") && (string(input[col][char+2]) == "S") {
					sum++
				}
				// M . M
				// . A .
				// S . S
				if (string(input[col][char]) == "M") && (string(input[col+1][char+1]) == "A") && (string(input[col+2][char+2]) == "S") && (string(input[col+2][char]) == "S") && (string(input[col][char+2]) == "M") {
					sum++
				}
				// S . S
				// . A .
				// M . M
				if (string(input[col][char]) == "S") && (string(input[col+1][char+1]) == "A") && (string(input[col+2][char+2]) == "M") && (string(input[col+2][char]) == "M") && (string(input[col][char+2]) == "S") {
					sum++
				}
				// S . M
				// . A .
				// S . M
				if (string(input[col][char]) == "S") && (string(input[col+1][char+1]) == "A") && (string(input[col+2][char+2]) == "M") && (string(input[col+2][char]) == "S") && (string(input[col][char+2]) == "M") {
					sum++
				}
			}

		}
	}

	fmt.Println(sum)
}
