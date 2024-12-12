package main

func invertString(str string) string {
	// Convert string to rune slice, reverse it, and return as a new string
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// How many times does needle or eldeen (inverted)
// occur in s.
func horizontalStringCount(needle, s string) int {
	count := 0
	for i := range s[:len(s)-len(needle)] {
		// Find needle
		if s[i:i+len(needle)] == needle {
			count++
			// Find eldeen (inverted)
		} else if s[i:i+len(needle)] == invertString(needle) {
			count++
		}
	}
	return count
}

func FindXmasWordCount(lines []string) int {
	xmasCount := 0

	for lineIndex, line := range lines {
		// Horizontal scans for the word XMAS in line
		xmasCount += horizontalStringCount("XMAS", line)
		for charIndex := range line {

			// Vertical scans
			if lineIndex < len(lines)-3 {
				if lines[lineIndex][charIndex] == 'X' &&
					lines[lineIndex+1][charIndex] == 'M' &&
					lines[lineIndex+2][charIndex] == 'A' &&
					lines[lineIndex+3][charIndex] == 'S' {
					xmasCount++
				} else if lines[lineIndex][charIndex] == 'S' &&
					lines[lineIndex+1][charIndex] == 'A' &&
					lines[lineIndex+2][charIndex] == 'M' &&
					lines[lineIndex+3][charIndex] == 'X' {
					xmasCount++
				}
			}

			// Diagonal scans
			if lineIndex < len(lines)-3 && charIndex < len(line)-3 {
				if lines[lineIndex][charIndex] == 'X' &&
					lines[lineIndex+1][charIndex+1] == 'M' &&
					lines[lineIndex+2][charIndex+2] == 'A' &&
					lines[lineIndex+3][charIndex+3] == 'S' {
					xmasCount++
				} else if lines[lineIndex][charIndex] == 'S' &&
					lines[lineIndex+1][charIndex+1] == 'A' &&
					lines[lineIndex+2][charIndex+2] == 'M' &&
					lines[lineIndex+3][charIndex+3] == 'X' {
					xmasCount++
				}
			}

			if lineIndex < len(lines)-3 && charIndex >= 3 {
				if lines[lineIndex][charIndex] == 'X' &&
					lines[lineIndex+1][charIndex-1] == 'M' &&
					lines[lineIndex+2][charIndex-2] == 'A' &&
					lines[lineIndex+3][charIndex-3] == 'S' {
					xmasCount++
				} else if lines[lineIndex][charIndex] == 'S' &&
					lines[lineIndex+1][charIndex-1] == 'A' &&
					lines[lineIndex+2][charIndex-2] == 'M' &&
					lines[lineIndex+3][charIndex-3] == 'X' {
					xmasCount++
				}
			}
		}

	}

	return xmasCount

}
