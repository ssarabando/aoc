package day07

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type nodeType int

const (
	Directory nodeType = iota
	File
)

type node struct {
	parent   *node
	name     string
	children nodes
	nodeType nodeType
	size     int
}

type nodes []*node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].size < n[j].size
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

var filesystem node

const (
	totalDiskSpace        int = 70_000_000
	requiredFreeDiskSpace int = 30_000_000
)

func fillFilesystem(lines []string) {
	filesystem = node{
		nil,
		"/",
		[]*node{},
		Directory,
		0,
	}

	currentDirectory := &filesystem

	for _, line := range lines {
		if line == "" {
			break
		}
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
				currentDirectory,
				newDirectoryName,
				[]*node{},
				Directory,
				0,
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
			currentDirectory,
			fileData[1],
			nil,
			File,
			size,
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

func PartOne(lines []string) int {
	fillFilesystem(lines)

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

func PartTwo(lines []string) int {
	fillFilesystem(lines)

	availableDiskSpace := totalDiskSpace - filesystem.size
	targetDiskSpace := requiredFreeDiskSpace - availableDiskSpace
	directories := getDirs(filesystem.children)
	sort.Sort(nodes(directories))
	for _, directory := range directories {
		if directory.size >= targetDiskSpace {
			return directory.size
		}
	}

	return filesystem.size
}
