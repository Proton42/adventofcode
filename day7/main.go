package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	List            = "$ ls"
	ChangeDirectory = "$ cd"
)

type Directory struct {
	FileSize int
	Name     string
	Parent   *Directory
	Children []*Directory
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	root := Directory{
		Name: "/",
	}
	currentDir := &root

	for scanner.Scan() {
		input := scanner.Text()

		if input == List {
			continue
		}

		if strings.HasPrefix(input, "$") { // directory command
			s := strings.Split(input, " ")

			if s[2] == ".." { // go up one directory
				currentDir = currentDir.Parent

				continue
			}

			nextDir := &Directory{
				Parent: currentDir,
				Name:   s[2],
			}

			currentDir.Children = append(currentDir.Children, nextDir)
			currentDir = nextDir

			continue // Process command
		}

		str := strings.Split(input, " ")

		if str[0] == "dir" {
			continue // Add directory
		}

		size, err := strconv.Atoi(str[0])
		if err != nil {
			log.Fatal(err)
		}

		currentDir.FileSize += size
	}

	total := 0
	goToDirectory(&root, &total)

	fmt.Println(total)
}

func goToDirectory(dir *Directory, total *int) {
	for _, child := range dir.Children {
		size := calcDirSize(child)

		if size <= 100000 {
			*total += size
		}

		goToDirectory(child, total)
	}
}

func calcDirSize(dir *Directory) int {
	childSizes := 0
	for _, child := range dir.Children {
		childSizes += calcDirSize(child)
	}

	return dir.FileSize + childSizes
}
