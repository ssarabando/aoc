package day06

import (
	"bufio"
	"log"
	"os"
)

const startOfPacket int = 4
const startOfMessage int = 14

// true if all chars in tail are different.
func isSignal(tail string) bool {
	chars := len(tail)
	for i := 0; i < chars; i++ {
		for j := i + 1; j < chars; j++ {
			if tail[i] == tail[j] {
				return false
			}
		}
	}
	return true
}

func processMessage(filename string, markerSize int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var message string
	var length int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		message += scanner.Text()
		length = len(message)
		if length < markerSize {
			continue
		}
		if isSignal(message[length-markerSize:]) {
			return length
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0

}

func PartOne(filename string) int {
	return processMessage(filename, startOfPacket)
}

func PartTwo(filename string) int {
	return processMessage(filename, startOfMessage)
}
