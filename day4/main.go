package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	overlapSum := 0

	for scanner.Scan() {
		if overlaps(scanner.Text()) {
			overlapSum += 1
		}
	}

	fmt.Println("Overlap Sum: ", overlapSum)
}

func overlaps(s string) bool {
	assignments := strings.Split(s, ",")
	fLower, fHigher := readRange(strings.Split(assignments[0], "-"))
	sLower, sHigher := readRange(strings.Split(assignments[1], "-"))

	return (fLower <= sLower && fHigher >= sHigher) || (sLower <= fLower && sHigher >= fHigher)
}

func readRange(sections []string) (int, int) {
	lower, _ := strconv.Atoi(sections[0])
	higher, _ := strconv.Atoi(sections[1])

	return lower, higher
}
