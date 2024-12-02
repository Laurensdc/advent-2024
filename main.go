package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var debugging = true

func main() {
	reports := readReports("day02_input.txt")
	output := countSafeReports(reports)

	fmt.Printf("%v\n", output)
}

// Sorts both lists and check the distance between each item
// Returns total distance between all items in the list
func calculateTotalDistance(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	totalDistance := 0

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}

		totalDistance += diff
	}

	return totalDistance
}

// A similarity score is calculated by adding up
// each number in the left list
// after multiplying it by the number of times that number appears in the right list.
func calculateSimilarityScore(list1, list2 []int) int {
	similarityScore := 0

	for i := 0; i < len(list1); i++ {
		// For each item in this list
		// Check how many times it exists in list2

		number := list1[i]
		numberFoundInList2 := 0

		for j := 0; j < len(list2); j++ {
			if list2[j] == number {
				numberFoundInList2++
			}
		}

		similarityScore += number * numberFoundInList2
	}

	return similarityScore
}

// Read reports from file and transform to [][]int
func readReports(filename string) [][]int {
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

func checkLevelsSafety(levels []int, withTolerance bool) bool {
	for i := 0; i < len(levels)-1; i++ {
		// The same number consecutively is not allowed
		if levels[i] == levels[i+1] {
			return false
		}

		// Have to keep increasing
		wasIncreasing := i > 1 && levels[i] > levels[i-1]
		if wasIncreasing && levels[i+1] < levels[i] {
			return false
		}

		// have to keep decreasing
		wasDecreasing := i > 1 && levels[i] < levels[i-1]
		if wasDecreasing && levels[i+1] > levels[i] {
			return false
		}

		diff := levels[i+1] - levels[i]
		if diff < 0 {
			diff = -diff
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	// No error cases met
	// means we kept increasing or decreasing
	// and the diff between 2 levels was always between 1 and 3
	return true
}

func countSafeReports(report [][]int) int {
	count := 0
	for i := 0; i < len(report); i++ {
		if checkLevelsSafety(report[i], false) == true {
			count++
		}
	}
	return count
}

func countSafeReportsWithTolerance(report [][]int) int {
	count := 0
	for i := 0; i < len(report); i++ {
		if checkLevelsSafety(report[i], true) == true {
			count++
		}
	}
	return count
}
