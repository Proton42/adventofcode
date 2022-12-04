package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open file
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var current, most, calories int

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		calories, _ = strconv.Atoi(line)
		most, current = process(most, current, calories)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("most caloreis: ", most)
}

func process(most, current, calories int) (int, int) {
	if calories == 0 {
		return most, 0
	}

	current += calories

	if current > most {
		most = current
	}

	return most, current
}
