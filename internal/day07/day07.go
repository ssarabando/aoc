package day07

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type nodeType int

const (
	Directory = iota
	File
)

type node struct {
	name     string
	nodeType nodeType
	size     int
	parent   *node
	children []*node
}

var input []string
var filesystem node

func readInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	input = []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func fillFilesystem() {
	filesystem = node{
		"/",
		Directory,
		0,
		nil,
		[]*node{},
	}

	currentDirectory := &filesystem

	for _, line := range input[1:] {
		if line[:4] == "$ cd" {
			if line[5] == '/' {
				currentDirectory = &filesystem
			} else if line[5] == '.' && line[5:7] == ".." {
				currentDirectory = currentDirectory.parent
			} else {
				targetDirectoryName := line[5:]
				for _, directory := range currentDirectory.children {
					if directory.name == targetDirectoryName {
						currentDirectory = directory
						break
					}
				}
			}
			continue
		}
		if line[:4] == "$ ls" {
			continue
		}
		if line[:3] == "dir" {
			newDirectoryName := line[4:]
			directory := &node{
				newDirectoryName,
				Directory,
				0,
				currentDirectory,
				[]*node{},
			}
			currentDirectory.children = append(
				currentDirectory.children,
				directory)
			continue
		}

		fileData := strings.Split(line, " ")
		size, err := strconv.Atoi(fileData[0])
		if err != nil {
			log.Fatal(err)
		}

		fileNode := &node{
			fileData[1],
			File,
			size,
			currentDirectory,
			nil,
		}
		currentDirectory.size += size
		currentDirectory.children = append(
			currentDirectory.children,
			fileNode)

		parent := currentDirectory.parent
		for parent != nil {
			parent.size += size
			parent = parent.parent
		}
	}
}

func getDirs(nodes []*node) []*node {
	result := []*node{}

	for _, node := range nodes {
		if node.nodeType != Directory {
			continue
		}
		result = append(result, node)
		result = append(result, getDirs(node.children)...)
	}

	return result
}

func PartOne(filename string) int {
	readInput(filename)
	fillFilesystem()

	result := 0

	currentDir := &filesystem
	directories := getDirs(currentDir.children)
	for _, directory := range directories {
		if directory.size <= 100_000 {
			result += directory.size
		}
	}

	return result
}
