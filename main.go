package main

import (
	"fmt"
	"slices"
)

func main() {
	output := calculateSimilarityScore(Day01_list1, Day01_list2)

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

func checkLevelsSafety(levels []int) bool {
	const (
		increasing = iota
		decreasing
	)

	var direction int

	if levels[1] > levels[0] {
		direction = increasing
	} else if levels[1] < levels[0] {
		direction = decreasing
	} else {
		// Neither increasing or decreasing means not safe
		return false
	}

	for i := 0; i < len(levels)-1; i++ {
		if direction == increasing {
			// Have to keep increasing or decreasing
			if levels[i+1] == levels[i] || levels[i+1] < levels[i] {
				return false
			}
		}

		if direction == decreasing {
			if levels[i+1] == levels[i] || levels[i+1] > levels[i] {
				return false
			}
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
		if checkLevelsSafety(report[i]) == true {
			count++
		}
	}
	return count
}
