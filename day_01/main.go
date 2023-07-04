package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var top1, top2, top3, current int
	var line string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			if current > top1 {
				top3 = top2
				top2 = top1
				top1 = current
			} else if current > top2 {
				top3 = top2
				top2 = current
			} else if current > top3 {
				top3 = current
			}
			current = 0
		} else {
			calories, _ := strconv.Atoi(line)
			current += calories
		}
	}

	fmt.Println(top1 + top2 + top3)
}
