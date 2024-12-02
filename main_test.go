package main

import (
	"testing"
)

func TestCalculateTotalDistance(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	output := calculateTotalDistance(list1, list2)
	expected := 11

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}

func TestSimilarityScore(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	output := calculateSimilarityScore(list1, list2)
	expected := 31

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}
