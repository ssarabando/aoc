package day11

import (
	"log"
	"strconv"
	"strings"
)

type bag struct {
	items []int
}

func newBag() *bag {
	return &bag{
		[]int{},
	}
}

func (b *bag) add(itemNumber int) {
	b.items = append(b.items, itemNumber)
}

func readStartingItems(line string) bag {
	bag := bag{}

	for _, item := range strings.Split(line[len(startingItemsLineKey):], ", ") {
		if itemNumber, err := strconv.Atoi(item); err != nil {
			log.Fatal(err)
		} else {
			bag.add(itemNumber)
		}
	}
	return bag
}
