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
		if contains(scanner.Text()) {
			overlapSum += 1
		}
	}

	fmt.Println("Overlap Sum: ", overlapSum)
}

func contains(s string) bool {
	assignments := strings.Split(s, ",")
	fSectionStart, fSectionEnd := sectionRange(strings.Split(assignments[0], "-"))
	sSectionStart, sSectionEnd := sectionRange(strings.Split(assignments[1], "-"))

	return (fSectionStart <= sSectionStart && fSectionEnd >= sSectionEnd) ||
		(sSectionStart <= fSectionStart && sSectionEnd >= fSectionEnd)
}

func sectionRange(sections []string) (int, int) {
	sectionStart, _ := strconv.Atoi(sections[0])
	sectionEnd, _ := strconv.Atoi(sections[1])

	return sectionStart, sectionEnd
}
