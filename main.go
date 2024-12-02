package main

import (
	"fmt"
	"slices"
)

func main() {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	output := getTotalDistance(list1, list2)

	fmt.Printf("%v\n", output)
}

func getTotalDistance(list1, list2 []int) int {
	sortList(list1)
	sortList(list2)

	return 5
}

func sortList(list []int) {
	slices.Sort(list)
}
