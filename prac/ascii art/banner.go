package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	result := make(map[rune][]string)

	for i := 0; i < 95; i++ {
		start := i * 9

		if start+8 > len(lines) {
			break
		}

		ch := rune(32 + i)

		art := make([]string, 8)
		copy(art, lines[start:start+8])

		result[ch] = art
	}

	return result, nil
}
