package main

import (
	"fmt"
)

func debugPrint(format string, a ...any) (n int, err error) {
	var debugging = false
	if debugging {
		return fmt.Printf(format+"\n", a...)
	}
	return 0, err
}

func main() {
	input := ReadXmasInput("day04_input.txt")
	output := FindXmasWordCount(input)

	fmt.Printf("%v\n", output)
}
