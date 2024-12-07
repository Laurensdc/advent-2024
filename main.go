package main

import (
	"fmt"
)

func debugPrint(format string, a ...any) (n int, err error) {
	var debugging = true
	if debugging {
		return fmt.Printf(format+"\n", a...)
	}
	return 0, err
}

func main() {
	reports := ReadReports("day02_input.txt")
	output := CountSafeReportsWithTolerance(reports)

	fmt.Printf("%v\n", output)
}
