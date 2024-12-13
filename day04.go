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
	for i := range s {
		// We will read len(needle) to the right,
		// don't go out of bounds
		if i > len(s)-len(needle) {
			continue
		}

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

// How many times does needle or eldeen (inverted)
// occur across lines (vertically downward)
func verticalStringCount(needle string, lines []string) int {
	count := 0

	for i := range lines {
		// We will read len(needle) down,
		// don't go out of bounds
		if i > len(lines)-len(needle) {
			continue
		}

		// Loop over every char in the line
		for charIndex := range lines[i] {

			// Get the word by fetching characters "down"
			// across lines
			var wordBytes []byte = make([]byte, len(needle))
			for n := range needle {
				wordBytes[n] = lines[i+n][charIndex]
			}

			word := string(wordBytes)

			if word == needle || word == invertString(needle) {
				count++
			}
		}
	}

	return count
}

func FindXmasWordCount(lines []string) int {
	xmasCount := 0

	xmasCount += verticalStringCount("XMAS", lines)

	for lineIndex, line := range lines {
		// Horizontal scans for the word XMAS in line
		xmasCount += horizontalStringCount("XMAS", line)

		for charIndex := range line {

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
