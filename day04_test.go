package main

import "testing"

func TestFindXmasWordCount(t *testing.T) {
	var input []string = []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	output := FindXmasWordCount(input)
	expected := 18

	if output != expected {
		t.Errorf("Expected\n%v\ngot\n%v\n", expected, output)
	}

}
