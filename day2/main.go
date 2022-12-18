package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissor
)

func ToMove(s string) (Move, error) {
	switch {
	case s == "A" || s == "X":
		return Rock, nil
	case s == "B" || s == "Y":
		return Paper, nil
	case s == "C" || s == "Z":
		return Scissor, nil
	default:
		return 0, errors.New("Oh no!")
	}
}

func ToResult(s string) (Result, error) {
	switch {
	case s == "X":
		return Loss, nil
	case s == "Y":
		return Draw, nil
	case s == "Z":
		return Win, nil
	default:
		return 0, errors.New("Oh no!")
	}
}

type Result int

const (
	Loss Result = 3 * iota
	Draw
	Win
)

func GetResult(theirs, yours Move) Result {
	if theirs == yours {
		return Draw
	}

	if theirs == Rock {
		if yours == Paper {
			return Win
		} else {
			return Loss
		}
	}

	if theirs == Paper {
		if yours == Scissor {
			return Win
		} else {
			return Loss
		}
	}

	if yours == Rock { // They have Scissors
		return Win
	}

	return Loss // We have paper
}

func yourMove(theirs Move, outcome Result) Move {
	if outcome == Draw {
		return theirs
	}

	if outcome == Win {
		if theirs == Paper {
			return Scissor
		}
		if theirs == Rock {
			return Paper
		}
		return Rock
	}

	if theirs == Paper {
		return Rock
	}
	if theirs == Rock {
		return Scissor
	}
	return Paper
}

func CalcPoints(theirs Move, outcome Result) int {
	yours := yourMove(theirs, outcome)
	res := int(GetResult(theirs, yours))

	result := res + int(yours)

	return result
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := 0

	for scanner.Scan() {
		theirs, err := ToMove(string(scanner.Text()[0]))
		if err != nil {
			log.Fatal(err)
		}

		yours, err := ToResult(string(scanner.Text()[2]))
		if err != nil {
			log.Fatal(err)
		}

		points += CalcPoints(theirs, yours)
	}

	fmt.Println("expected points", points)
}
