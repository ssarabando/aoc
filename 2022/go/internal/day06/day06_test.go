package day06

import (
	"log"
	"os"
	"strings"
	"testing"
)

func getContents(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\r\n")
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, test := range tests {
		data := test.data
		actual := PartOne([]string{data})
		expected := test.expected
		if actual != expected {
			t.Fatalf(`%s: expected %d; found %d`, data, expected, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, test := range tests {
		data := test.data
		actual := PartTwo([]string{data})
		expected := test.expected
		if actual != expected {
			t.Fatalf(`%s: expected %d; found %d`, data, expected, actual)
		}
	}
}
