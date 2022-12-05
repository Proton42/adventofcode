package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Priority(r int) int {
	if r <= 90 {
		return r - 38
	}
	return r - 96
}

func exists(item rune, compartment string) bool {
	for _, s := range compartment {
		if s == item {
			return true
		}
	}

	return false
}

func SharedItems(comp1, comp2 string) []rune {
	itemMap := map[rune]bool{}

	for _, s := range comp1 {
		if exists(s, comp2) {
			itemMap[s] = true
		}
	}

	items := []rune{}
	for k, _ := range itemMap {
		items = append(items, k)
	}

	return items
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	prioritySum := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		sharedItems := SharedItems(compartment1, compartment2)

		for _, item := range sharedItems {
			prioritySum += Priority(int(item))
		}
	}

	fmt.Println("Priority Sum: ", prioritySum)
}
