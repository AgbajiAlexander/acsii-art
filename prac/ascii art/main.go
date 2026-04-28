package main

import (
	"fmt"
	"log"
)

func main() {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total characters loaded:", len(banner))

	fmt.Println("\nASCII art for A:")
	fmt.Println("----------------")

	for _, line := range banner['A'] {
		fmt.Println(line)
	}
}
