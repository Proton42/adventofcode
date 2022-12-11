package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()

	for i := 0; i <= len(input); i++ {
		if isMarker(input[i : i+4]) {
			fmt.Println("Start of package: ", i+4)
			break
		}
	}
}

func isMarker(chars string) bool {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if chars[i] == chars[j] {
				return false
			}
		}
	}

	return true
}
