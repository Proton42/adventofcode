package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Right Direction = "R"
	Left  Direction = "L"
	Down  Direction = "D"
	Up    Direction = "U"
)

type Knot struct {
	x int
	y int
}

type Rope struct {
	Knots [10]Knot
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rope := Rope{}

	marker := [1000][1000]bool{}
	print(rope)
	for scanner.Scan() {
		direction, moves := readLine(scanner.Text())

		for i := 0; i < moves; i++ {
			moveRope(direction, &rope)

			marker[rope.Knots[len(rope.Knots)-1].x+500][rope.Knots[len(rope.Knots)-1].y+500] = true
		}
	}

	count := 0
	for i := 0; i < len(marker); i++ {
		for j := 0; j < len(marker[i]); j++ {
			if marker[i][j] {
				count++
			}
		}
	}

	fmt.Println(count)
}

func moveRope(dir Direction, rope *Rope) {
	moveHead(dir, &rope.Knots[0])

	for i := 1; i < len(rope.Knots); i++ {
		moved := moveKnot(&rope.Knots[i-1], &rope.Knots[i])

		if !moved {
			return
		}
	}
}

func moveHead(dir Direction, k *Knot) {
	switch dir {
	case Right:
		k.x++
	case Left:
		k.x--
	case Up:
		k.y++
	case Down:
		k.y--
	}
}

func moveKnot(fKnot, sKnot *Knot) bool {
	if abs(fKnot.x-sKnot.x) <= 1 && abs(fKnot.y-sKnot.y) <= 1 {
		return false
	}

	if fKnot.x != sKnot.x && fKnot.y != sKnot.y {
		if fKnot.x > sKnot.x && fKnot.y > sKnot.y { // top right
			sKnot.x++
			sKnot.y++
		}

		if fKnot.x > sKnot.x && fKnot.y < sKnot.y { // down right
			sKnot.x++
			sKnot.y--
		}

		if fKnot.x < sKnot.x && fKnot.y > sKnot.y { // top left
			sKnot.x--
			sKnot.y++
		}

		if fKnot.x < sKnot.x && fKnot.y < sKnot.y { // down right
			sKnot.x--
			sKnot.y--
		}

		return true
	}

	if fKnot.x > sKnot.x { // right
		sKnot.x++
	} else if fKnot.x < sKnot.x { // left
		sKnot.x--
	} else if fKnot.y > sKnot.y { // above
		sKnot.y++
	} else if fKnot.y < sKnot.y { // below
		sKnot.y--
	}

	return true
}

func readLine(line string) (Direction, int) {
	s := strings.Split(line, " ")

	direction := Direction(s[0])
	moves, _ := strconv.Atoi(s[1])

	return direction, moves
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

var debug = 0

func print(rope Rope) {
	offset := 16
	fmt.Println()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Debug: ", debug)
	debug++

	for i := offset*2 - 1; i > 0; i-- {
		fmt.Println(i + 1)

		for j := 0; j < offset*2; j++ {
			symbol := "."
			for y := 1; y < len(rope.Knots); y++ {
				if rope.Knots[y].x == j-offset && rope.Knots[y].y == i-offset {
					symbol = strconv.Itoa(y)
					break
				}
			}

			if rope.Knots[0].x == j-offset && rope.Knots[0].y == i-offset {
				symbol = "H"
			}

			fmt.Print(symbol)
		}
	}
	fmt.Println()
}
