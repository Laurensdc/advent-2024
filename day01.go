package main

import "slices"

// Sorts both lists and check the distance between each item
// Returns total distance between all items in the list
func CalculateTotalDistance(list1, list2 []int) int {
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
func CalculateSimilarityScore(list1, list2 []int) int {
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
