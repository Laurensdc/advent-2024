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
		} else {
			cleanedUpReport := make([]int, len(report[i]))
			copy(cleanedUpReport, report[i])

			// Remove item at :badLevelIndex and retry
			cleanedUpReport = append(cleanedUpReport[:badLevelIndex], cleanedUpReport[badLevelIndex+1])
			badLevelIndex = getBadLevelIndex(cleanedUpReport)

			if badLevelIndex == -1 {
				count++
			}

		}
	}
	return count
}

// Check if we are ascending or descending
// and take into account the tolerance of 1 odd number.
// As soon as 2 items are ascending or descending, we are certain
// of the direction
func getDirection(levels []int) int {
	direction := 0
	ascendingCount := 0
	descendingCount := 0
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] < levels[i]+1 {
			ascendingCount++
			if ascendingCount >= 2 {
				direction = ascending
				break
			}

		}
		if levels[i] > levels[i+1] {
			descendingCount++
			if descendingCount >= 2 {
				direction = descending
				break
			}
		}
	}

	if debugging {
		fmt.Printf("Direction: %v\n", direction)
	}

	return direction
}

// -1 means all levels were good
func getBadLevelIndex(levels []int) int {
	if debugging {
		fmt.Printf("\n*** Checking levels %v\n", levels)
	}

	direction := getDirection(levels)

	for i := 0; i < len(levels)-1; i++ {
		if debugging {
			fmt.Printf("Processing %v and %v for i=%v\n",
				levels[i], levels[i+1], i)
		}

		if levels[i] == levels[i+1] {
			fmt.Printf("NOT SAFE, %v == %v at %v\n", levels[i], levels[i+1], i)
			return i

		}

		if direction == ascending && levels[i] > levels[i+1] {
			fmt.Printf("NOT SAFE, err in asc ordering for %v, %v at i=%v\n", levels[i], levels[i+1], i)
			return i
		}

		if direction == descending && levels[i] < levels[i+1] {
			fmt.Printf("NOT SAFE, err in desc ordering for %v, %v at i=%v\n", levels[i], levels[i+1], i)
			return i
		}

		diff := levels[i+1] - levels[i]
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			fmt.Printf("NOT SAFE, diff is %v at i=%v\n", diff, i)

		}
	}
	if debugging {
		fmt.Println("SAFE")
	}

	return -1
}

// TODO: Delete me?
func removeItemFromSlice(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
