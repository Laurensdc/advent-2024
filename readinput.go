package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read reports from file and transform to [][]int
func ReadReports(filename string) [][]int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Failed to read file %v because of %v\n", filename, err)
		os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var reports = [][]int{}
	for scanner.Scan() {
		var numbers []int

		numbersAsStrings := strings.Fields(scanner.Text())
		for i := 0; i < len(numbersAsStrings); i++ {
			number, err := strconv.Atoi(numbersAsStrings[i])
			if err != nil {
				fmt.Printf("Failed to parse str %v to int because of %v\n", numbersAsStrings[i], err)
				os.Exit(1)
			}
			numbers = append(numbers, number)
		}

		reports = append(reports, numbers)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("Failed to read file %v because of %v\n", filename, err)
		os.Exit(1)
	}

	return reports
}
