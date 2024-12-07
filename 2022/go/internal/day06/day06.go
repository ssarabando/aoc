package day06

const (
	startOfPacket  int = 4
	startOfMessage int = 14
)

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

func processMessage(lines []string, markerSize int) int {
	var message string
	var length int

	for _, line := range lines {
		for _, rune := range line {
			message += string(rune)
			length = len(message)
			if length < markerSize {
				continue
			}
			if isSignal(message[length-markerSize:]) {
				return length
			}
		}
	}

	return 0
}

func PartOne(lines []string) int {
	return processMessage(lines, startOfPacket)
}

func PartTwo(lines []string) int {
	return processMessage(lines, startOfMessage)
}
