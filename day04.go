package main

func FindXmasWordCount(lines []string) int {
	xmasCount := 0

	xmasCount += horizontalStringCount("XMAS", lines)
	xmasCount += verticalStringCount("XMAS", lines)
	xmasCount += diagonalStringCount("XMAS", lines)

	return xmasCount
}

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
func horizontalStringCount(needle string, lines []string) int {
	count := 0
	for _, s := range lines {
		// Horizontal scans for the word XMAS in line

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

// How many times does needle or eldeen (inverted)
// occur diagonally across lines (vertically downward)
func diagonalStringCount(needle string, lines []string) int {
	count := 0

	for i := range lines {
		// We will read len(needle) down,
		// don't go out of bounds
		if i > len(lines)-len(needle) {
			continue
		}

		// Loop over every char in the line
		for charIndex := range lines[i] {
			// Get the word by fetching characters diagonally down and to the right
			if charIndex+len(needle) <= len(lines[i]) {
				var wordBytes []byte = make([]byte, len(needle))
				for n := range needle {
					wordBytes[n] = lines[i+n][charIndex+n]
				}
				word := string(wordBytes)

				if word == needle || word == invertString(needle) {
					count++
				}
			}

			// Get the word by fetching characters diagonally down and to the left
			if charIndex+1-len(needle) >= 0 {
				var wordBytes []byte = make([]byte, len(needle))
				for n := range needle {
					wordBytes[n] = lines[i+n][charIndex-n]
				}
				word := string(wordBytes)

				if word == needle || word == invertString(needle) {
					count++
				}
			}
		}
	}

	return count
}
