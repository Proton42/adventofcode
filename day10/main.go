package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func noop(x int) int {
	if DEBUG {
		fmt.Println("noop")
	}
	return x
}

func addx(input int) func(int) int {
	return func(x int) int {
		if DEBUG {
			fmt.Println(input)
		}
		return x + input
	}
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	registerX := 1

	operations := []func(int) int{}
	for scanner.Scan() {
		operations = append(operations, read(scanner.Text())...)
	}

	signalSum := 0
	for i, op := range operations {
		cycle := i + 1
		if (cycle+20)%40 == 0 {
			// fmt.Println("Cycle: ", cycle, " , Value: ", registerX, ", Strength: ", registerX*cycle)
			signalSum += cycle * registerX
		}

		if i%40 == 0 {
			fmt.Println(" i: ", i)
		}

		drawPixel(i%40, registerX)

		registerX = op(registerX)
	}

	fmt.Println("Signal strength sum: ", signalSum)
}

func drawPixel(position int, x int) {
	if abs(position-x) > 1 {
		fmt.Print(".")
		return
	}

	fmt.Print("#")
}

func read(input string) []func(int) int {
	if input == "" { // end of input
		return []func(int) int{}
	}

	s := strings.Split(input, " ")
	ops := []func(int) int{noop}

	if s[0] == "noop" {
		return ops // only return noop
	}

	value, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}

	return append(ops, addx(value)) // return noop and addx
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
