package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func cleanCorruptedMemory(s string) []string {
	// Valid syntax is
	// "mul("
	// 1-3 digits
	// ","
	// 1-3 digits
	// ")"
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := re.FindAllString(s, -1)

	debugPrint("String matches %v", matches)

	return matches
}

func calculateInstruction(s string) int {
	// Remove "mul("
	str := s[4:]

	// Extract numbers
	numbers := strings.Split(str, ",")
	n1Str := numbers[0]
	n2Str := numbers[1][:len(numbers[1])-1] // Chop off the last ")"

	debugPrint("Extracted numbers: %v & %v", n1Str, n2Str)

	n1, err := strconv.Atoi(n1Str)
	if err != nil {
		fmt.Printf("Error parsing value from %v\n%v\n", n1Str, err)
		os.Exit(-1)
	}
	n2, err := strconv.Atoi(n2Str)
	if err != nil {
		fmt.Printf("Error parsing value from %v\n%v\n", n1Str, err)
		os.Exit(-1)
	}

	return n1 * n2
}

func CalculateFromCorruptedMemory(s string) int {
	slices := cleanCorruptedMemory(s)
	total := 0
	for i := range slices {
		total += calculateInstruction(slices[i])
	}
	return total
}
