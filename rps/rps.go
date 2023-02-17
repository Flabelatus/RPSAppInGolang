package rps

import (
	"math/rand"
	"time"
)

var PlayerWinMsg = []string{
	"Welldone",
	"Perfect",
	"Great Job You Won!",
	"WOHOOO!",
}
var ComputerWinMsg = []string{
	"Damn!",
	"Huh! Try Again",
	"Too Bad!",
	"...!",
}
var DrawWinMsg = []string{
	"HAha! one more time",
	"Hmm!",
	"Ok try again",
	"Atleast you dont loose",
}

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	Message        string `json:"message"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)

	computerChoice := ""
	roundResult := ""
	message := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
		message = DrawWinMsg[rand.Intn(len(DrawWinMsg))]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS
		message = PlayerWinMsg[rand.Intn(len(PlayerWinMsg))]
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
		message = ComputerWinMsg[rand.Intn(len(ComputerWinMsg))]
	}

	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	result.Message = message

	return result
}
