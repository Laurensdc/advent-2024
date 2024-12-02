package main

import (
	"fmt"
	"slices"
)

func main() {
	output := getTotalDistance(Ex1_list1, Ex1_list2)

	fmt.Printf("%v\n", output)
}

func getTotalDistance(list1, list2 []int) int {
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

func sortList(list []int) {
}
