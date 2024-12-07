package day02

import (
	"strings"
)

type Shape struct {
	name string
}

type Play struct {
	opponent, mine Shape
}

const (
	score_lost int = 0
	score_draw int = 3
	score_won  int = 6
)

func PartTwo(lines []string) int {
	rock := Shape{"Rock"}
	paper := Shape{"Paper"}
	scissors := Shape{"Scissors"}

	game_scores := map[Play]int{
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

	losing_plays := map[Shape]Shape{
		rock:     scissors,
		scissors: paper,
		paper:    rock,
	}

	draw_plays := map[Shape]Shape{
		rock:     rock,
		scissors: scissors,
		paper:    paper,
	}

	winning_plays := map[Shape]Shape{
		rock:     paper,
		paper:    scissors,
		scissors: rock,
	}

	shape_scores := map[Shape]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}

	opponent_shapes := map[rune]Shape{
		'A': rock,
		'B': paper,
		'C': scissors,
	}

	var total_score int

	for _, line := range lines {
		if line == "" {
			break
		}

		current_play := strings.Fields(line)

		opponent_play, my_play := rune(current_play[0][0]), rune(current_play[1][0])

		opponent_shape := opponent_shapes[opponent_play]

		var my_shape Shape
		switch my_play {
		case 'X':
			my_shape = losing_plays[opponent_shape]
		case 'Y':
			my_shape = draw_plays[opponent_shape]
		default:
			my_shape = winning_plays[opponent_shape]
		}

		outcome := shape_scores[my_shape] + game_scores[Play{opponent_shape, my_shape}]

		total_score += outcome
	}

	return total_score
}
