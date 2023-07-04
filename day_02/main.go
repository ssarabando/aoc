package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape struct {
	name string
}

type Play struct {
	opponent, mine Shape
}

const score_lost int = 0
const score_draw int = 3
const score_won int = 6

func main() {
	var rock = Shape{"Rock"}
	var paper = Shape{"Paper"}
	var scissors = Shape{"Scissors"}

	var outcomes = map[Play]int{
		{rock, scissors}:     score_lost,
		{scissors, paper}:    score_lost,
		{paper, rock}:        score_lost,
		{rock, rock}:         score_draw,
		{paper, paper}:       score_draw,
		{scissors, scissors}: score_draw,
		{scissors, rock}:     score_won,
		{paper, scissors}:    score_won,
		{rock, paper}:        score_won,
	}

	var shape_scores = map[Shape]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}

	var opponent_shapes = map[rune]Shape{
		'A': rock,
		'B': paper,
		'C': scissors,
	}

	var my_shapes = map[rune]Shape{
		'X': rock,
		'Y': paper,
		'Z': scissors,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total_score int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current_play := strings.Fields(scanner.Text())

		opponent_play, my_play := rune(current_play[0][0]), rune(current_play[1][0])

		opponent_shape := opponent_shapes[opponent_play]

		my_shape := my_shapes[my_play]

		outcome := shape_scores[my_shape] + outcomes[Play{opponent_shape, my_shape}]

		fmt.Printf(
			"Opponent: %s (%s); me: %s (%s) = %d\n",
			string(opponent_play), opponent_shape.name,
			string(my_play), my_shape.name,
			outcome)

		total_score += outcome
	}

	fmt.Printf("Total score: %d", total_score)
}
