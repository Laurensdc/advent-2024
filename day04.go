package main

import ()

func FindXmasWordCount(lines []string) int {
	xmasCount := 0

	for lineIndex, line := range lines {
		for charIndex := range line {
			// Horizontal scans for the word XMAS in line
			if charIndex < len(line)-3 {
				// if line[charIndex:charIndex+4] == "XMAS" {
				// 	xmasCount++
				// } else if line[charIndex:charIndex+4] == "SAMX" {
				// 	xmasCount++
				// }

				if line[charIndex] == 'X' &&
					line[charIndex+1] == 'M' &&
					line[charIndex+2] == 'A' &&
					line[charIndex+3] == 'S' {
					xmasCount++
				} else if line[charIndex] == 'S' &&
					line[charIndex+1] == 'A' &&
					line[charIndex+2] == 'M' &&
					line[charIndex+3] == 'X' {
					xmasCount++
				}

			}

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
