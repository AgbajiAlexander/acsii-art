package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Each char = 8 lines + 1 separator = 9 lines total
	// 95 printable ASCII chars from 32 to 126
	const charHeight = 8
	const charCount = 95
	expectedLines := charCount * (charHeight + 1)

	if len(lines) < expectedLines {
		return nil, fmt.Errorf("invalid banner file: expected at least %d lines, got %d", expectedLines, len(lines))
	}

	bannerMap := make(map[rune][]string)

	for i := 0; i < charCount; i++ {
		start := i * (charHeight + 1)
		end := start + charHeight

		if end > len(lines) {
			return nil, fmt.Errorf("unexpected end of file at character %d", i)
		}

		// ASCII code starts at 32 for space
		char := rune(32 + i)
		bannerMap[char] = lines[start:end]

		// Skip the separator line between characters
		// Last character doesn't have a separator after it
		if i < charCount-1 && start+charHeight < len(lines) {
			sepLine := lines[start+charHeight]
			if sepLine != "" {
				return nil, fmt.Errorf("expected blank separator line at %d, got %q", start+charHeight, sepLine)
			}
		}
	}

	return bannerMap, nil
}

func main() {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Test: print the letter 'A' which is ASCII 65
	fmt.Println("Character 'A':")
	for _, line := range banner['A'] {
		fmt.Println(line)
	}

	// Test: print space character
	fmt.Println("\nCharacter ' ' (space):")
	for _, line := range banner[' '] {
		fmt.Printf("%q\n", line) // Use %q to see empty strings clearly
	}

	fmt.Printf("\nLoaded %d characters\n", len(banner))
}
