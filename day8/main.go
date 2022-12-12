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

	forest := mapForest(bufio.NewScanner(f))

	count := countVisibleTrees(forest)

	fmt.Println(count)
}

func countVisibleTrees(forest [][]int) int {
	count := 0

	for y := 1; y < len(forest)-1; y++ {
		for x := 1; x < len(forest[y])-1; x++ {
			if treeIsVisible(y, x, forest) {
				count++
			}
		}
	}

	length := len(forest)

	return count + length*4 - 4
}

//     |
//     v
// 3 0 3 7 3
// 2 5 5 1 2
// 6 5 3 3 2
// 3 3 5 4 9 <---
// 3 5 3 9 0

// X = 2, Y = 3

func treeIsVisible(x, y int, forest [][]int) bool {
	east, north, west, south := true, true, true, true

	for i := 0; i < y; i++ {
		if forest[x][i] >= forest[x][y] {
			north = false

			break
		}
	}

	for i := y + 1; i < len(forest); i++ {
		if forest[x][i] >= forest[x][y] {
			south = false

			break
		}
	}

	for i := 0; i < x; i++ {
		if forest[i][y] >= forest[x][y] {
			west = false

			break
		}
	}

	for i := x + 1; i < len(forest[x]); i++ {
		if forest[i][y] >= forest[x][y] {
			east = false

			break
		}
	}

	return east || north || west || south
}

func mapForest(scanner *bufio.Scanner) [][]int {
	forest := [][]int{}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()

		forest = append(forest, []int{})

		for _, tree := range row {
			height, err := strconv.Atoi(string(tree))
			if err != nil {
				log.Fatal(err)
			}

			forest[i] = append(forest[i], height)
		}
	}

	return forest
}
