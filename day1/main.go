package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var currentCal, mostCal, calLine int
	for scanner.Scan() {
		calLine, _ = strconv.Atoi(scanner.Text())
		mostCal, currentCal = process(mostCal, currentCal, calLine)
	}

	fmt.Println("most calories", mostCal)
}

func process(most, current, calories int) (int, int) {
	if calories == 0 {
		return most, 0
	}

	if current += calories; current > most {
		most = current
	}

	return most, current
}
