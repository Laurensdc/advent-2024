package main

import "fmt"

func CountSafeReports(report [][]int) int {
	count := 0
	for i := 0; i < len(report); i++ {
		if getBadLevelIndex(report[i]) == -1 {
			count++
		}
	}
	return count
}

func CountSafeReportsWithTolerance(report [][]int) int {
	count := 0
	for i := 0; i < len(report); i++ {
		badLevelIndex := getBadLevelIndex(report[i])
		if debugging {
			fmt.Printf("Bad level index: %v\n", badLevelIndex)
		}

		if badLevelIndex == -1 { // Good level
			if debugging {
				fmt.Println("REPORT SAFE")
			}
			count++
			continue
		}

		// First retry
		cleanedUpReport := make([]int, len(report[i]))
		copy(cleanedUpReport, report[i])
		cleanedUpReport = removeItemFromSlice(cleanedUpReport, badLevelIndex)

		if debugging {
			fmt.Printf("CleanedUpReport %v\n", cleanedUpReport)
		}
		badLevelIndex2 := getBadLevelIndex(cleanedUpReport)

		if badLevelIndex2 == -1 {
			if debugging {
				fmt.Println("REPORT SAFE")
			}
			count++
			continue

		}

		// Second retry
		cleanedUpReport2 := make([]int, len(report[i]))
		copy(cleanedUpReport2, report[i])
		cleanedUpReport2 = removeItemFromSlice(cleanedUpReport2, badLevelIndex+1)

		if debugging {
			fmt.Printf("CleanedUpReport2 %v\n", cleanedUpReport2)
		}
		badLevelIndex = getBadLevelIndex(cleanedUpReport2)

		if badLevelIndex == -1 {
			if debugging {
				fmt.Println("REPORT SAFE")
			}
			count++
			continue
		}
		if debugging {
			fmt.Println("REPORT NOT SAFE")
		}

	}
	return count
}

// Check if we are ascending or descending
// and take into account the tolerance of 1 odd number.
// As soon as 2 items are ascending or descending, we are certain
// of the direction
func isAscending(levels []int) bool {
	ascendingCount := 0
	descendingCount := 0

	for i := 0; i < len(levels)-1; i++ {
		if levels[i] < levels[i+1] {
			ascendingCount++
			if ascendingCount >= 2 {
				return true
			}

		}

		if levels[i] > levels[i+1] {
			descendingCount++
			if descendingCount >= 2 {
				return false
			}
		}
	}

	fmt.Println("Error! Not ascending nor descending")
	return false
}

// -1 means all levels were good
func getBadLevelIndex(levels []int) int {
	if debugging {
		fmt.Printf("\n*** Checking levels %v\n", levels)
	}

	isAscending := isAscending(levels)

	for i := 0; i < len(levels)-1; i++ {
		if debugging {
			fmt.Printf("Processing %v and %v for i=%v\n",
				levels[i], levels[i+1], i)
		}

		if levels[i] == levels[i+1] {
			fmt.Printf("NOT SAFE, %v == %v at %v\n", levels[i], levels[i+1], i)
			return i

		}

		if isAscending == true && levels[i] > levels[i+1] {
			fmt.Printf("NOT SAFE, err in asc ordering for %v, %v at i=%v\n", levels[i], levels[i+1], i)
			return i
		}

		if isAscending == false && levels[i] < levels[i+1] {
			fmt.Printf("NOT SAFE, err in desc ordering for %v, %v at i=%v\n", levels[i], levels[i+1], i)
			return i
		}

		diff := levels[i+1] - levels[i]
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			fmt.Printf("NOT SAFE, diff is %v at i=%v\n", diff, i)
			return i

		}
	}

	if debugging {
		fmt.Println("SAFE")
	}

	return -1
}

func removeItemFromSlice(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
