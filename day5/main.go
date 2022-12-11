package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stacks = map[int][]rune{
	1: {'Q', 'S', 'W', 'C', 'Z', 'V', 'F', 'T'},
	2: {'Q', 'R', 'B'},
	3: {'B', 'Z', 'T', 'Q', 'P', 'M', 'S'},
	4: {'D', 'V', 'F', 'R', 'Q', 'H'},
	5: {'J', 'G', 'L', 'D', 'B', 'S', 'T', 'P'},
	6: {'W', 'R', 'T', 'Z'},
	7: {'H', 'Q', 'M', 'N', 'S', 'F', 'R', 'J'},
	8: {'R', 'N', 'F', 'H', 'W'},
	9: {'J', 'Z', 'T', 'Q', 'P', 'R', 'B'},
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		from, to, nbrMoves := parseMove(scanner.Text())
		move(from, to, nbrMoves)
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Printf("Stack %d: %s\n", i, string(stacks[i][len(stacks[i])-1]))
	}
}

func parseMove(move string) (from, to, nbrMoves int) {
	splt := strings.Split(move, " ")

	nbrMoves, _ = strconv.Atoi(splt[1])
	from, _ = strconv.Atoi(splt[3])
	to, _ = strconv.Atoi(splt[5])

	return
}

func move(from, to, nbrMoves int) {
	var crate rune
	for i := 0; i < nbrMoves; i++ {
		crate, stacks[from] = stacks[from][len(stacks[from])-1], stacks[from][:len(stacks[from])-1]

		stacks[to] = append(stacks[to], crate)
	}
}
