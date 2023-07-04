package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const priorites string = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum_of_priorities := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		number_of_items := len(line)

		mid_point := number_of_items / 2

		for pos := 0; pos < mid_point; pos++ {
			item := string(line[pos])
			if strings.Contains(line[mid_point:number_of_items], item) {
				sum_of_priorities += strings.Index(priorites, item)
				break
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of priorites:", sum_of_priorities)
}
