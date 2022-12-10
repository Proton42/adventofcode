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

	var currentCal, calLine int

	ok := false
	top := [3]int{}

	for scanner.Scan() {
		calLine, _ = strconv.Atoi(scanner.Text())
		currentCal, ok = process(currentCal, calLine)

		if ok {
			top = checkIfMore(top, currentCal)
			currentCal = 0
		}
	}

	fmt.Println("most calories: %d", top[0]+top[1]+top[2])
}

func process(total, calories int) (int, bool) {
	if calories == 0 {
		return total, true
	}

	return total + calories, false
}

func checkIfMore(top [3]int, current int) [3]int {
	if current <= top[2] { // Not top 3
		return top
	}

	if current <= top[1] { // Top 3
		return [3]int{top[0], top[1], current}
	}

	if current <= top[0] { // Top 2
		return [3]int{top[0], current, top[1]}
	}

	return [3]int{current, top[0], top[1]} // Top 1
}
