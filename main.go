package main

import (
	"fmt"
)

var debugging = true

const (
	ascending = iota
	descending
)

func main() {
	reports := ReadReports("day02_input.txt")
	output := CountSafeReportsWithTolerance(reports)

	fmt.Printf("%v\n", output)
}
