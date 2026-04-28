package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter only 2 arguments!")
		return
	}

	input := os.Args[1]
	fmt.Println(AsciiArt(input))
}

func AsciiArt(input string) string {
	inputFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	inputFileLines := strings.Split(string(inputFile), "\n")
	words := strings.Split(input, "\\n")
	result := ""

	for _, word := range words {
		for i := 0; i < 8; i++ {
			for _, char := range word {
				result += inputFileLines[i+(int(char-' ')*9)+1]
			}
			result += "\n"
		}

	}
	return result
}
