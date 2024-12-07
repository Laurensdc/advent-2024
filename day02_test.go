package main

import (
	"fmt"
	"testing"
)

func TestGetDirection(t *testing.T) {
	input := []int{1, 5, 13, 24}
	output := isAscending(input)
	expected := true
	if output != expected {
		t.Errorf("Testing %v\nExpected\n%v\ngot\n%v\n", input, expected, output)
	}

	input = []int{100, 99, 80, 24, 5, 1}
	fmt.Println(input)
	output = isAscending(input)
	expected = false
	if output != expected {
		t.Errorf("Testing %v\nExpected\n%v\ngot\n%v\n", input, expected, output)
	}

	input = []int{10, 9, 8, 10, 7, 6}
	output = isAscending(input)
	expected = false
	if output != expected {
		t.Errorf("Testing %v\nExpected\n%v\ngot\n%v\n", input, expected, output)
	}

	input = []int{1, 3, 6, 7, 6, 9, 10}
	output = isAscending(input)
	expected = true
	if output != expected {
		t.Errorf("Testing %v\nExpected\n%v\ngot\n%v\n", input, expected, output)
	}

}

func TestReportSafety(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := 2
	output := CountSafeReports(reports)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}

func TestReportSafetyWithToleranceBasic(t *testing.T) {
	t.Skip()
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := 4
	output := CountSafeReportsWithTolerance(reports)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}

func TestReportSafetyWithToleranceEdgeCases(t *testing.T) {
	t.Skip()
	reports := [][]int{
		{48, 46, 47, 49, 51, 54, 56},
		{1, 1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 5},
		{5, 1, 2, 3, 4, 5},
		{1, 4, 3, 2, 1},
		{1, 6, 7, 8, 9},
		{1, 2, 3, 4, 3},
		{9, 8, 7, 6, 7},
		{7, 10, 8, 10, 11},
		{29, 28, 27, 25, 26, 25, 22, 20},
		// {24, 25, 28, 31, 28},           // safe if we remove the last 28
		// {41, 44, 45, 48, 49, 50, 50},   // safe, 1 bad level
		// {5, 8, 10, 13, 15, 16, 17, 21}, // safe, 1 bad level
		// {11, 13, 16, 17, 19, 26},
		// {79, 81, 78, 79, 82, 84},
		// {16, 19, 20, 18, 20, 22, 25, 22},
		// {84, 87, 90, 92, 94, 97, 96, 96},
		// {86, 87, 88, 91, 88, 91, 95},
		// {40, 43, 41, 44, 49},
		// {8, 10, 10, 11, 13},
	}
	expected := 10
	output := CountSafeReportsWithTolerance(reports)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}
