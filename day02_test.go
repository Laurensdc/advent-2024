package main

import (
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
		{7, 6, 4, 2, 1}, // safe
		{1, 2, 7, 8, 9}, // unsafe, gap > 3
		{9, 7, 6, 2, 1}, // unsafe, gap > 3
		{1, 3, 2, 4, 5}, // unsafe, doesnt keep ascending
		{8, 6, 4, 4, 1}, // unsafe, 2 equal levels
		{1, 3, 6, 7, 9}, // safe
	}
	expected := 2
	output := CountSafeReports(reports)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}

func TestReportSafetyWithToleranceBasic(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1}, // safe as is
		{1, 2, 7, 8, 9}, // unsafe
		{9, 7, 6, 2, 1}, // unsafe
		{1, 3, 2, 4, 5}, // safe by removing the 3
		{8, 6, 4, 4, 1}, // safe by removing the 4
		{1, 3, 6, 7, 9}, // safe as is
	}
	expected := 4
	output := CountSafeReportsWithTolerance(reports)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}

func TestReportSafetyWithToleranceEdgeCases(t *testing.T) {
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
	}
	expected := 10
	output := CountSafeReportsWithTolerance(reports)

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}
}
