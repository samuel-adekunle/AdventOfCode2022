package main

import (
	"fmt"
	"os"
	"strings"
)

type Move int

const (
	Unknown Move = iota
	Rock
	Paper
	Scissors
)

func getMove(move string) Move {
	if move == "A" {
		return Rock
	}
	if move == "B" {
		return Paper
	}
	if move == "C" {
		return Scissors
	}
	return Unknown
}

type Outcome int

const (
	Undefined Outcome = -1
	Loss      Outcome = 0
	Draw      Outcome = 3
	Win       Outcome = 6
)

func getOutcome(outcome string) Outcome {
	if outcome == "X" {
		return Loss
	}
	if outcome == "Y" {
		return Draw
	}
	if outcome == "Z" {
		return Win
	}
	return Undefined
}

func calculateOutcome(you, them Move) Outcome {
	if you == them {
		return Draw
	}
	if you == Rock && them == Scissors {
		return Win
	}
	if you == Paper && them == Rock {
		return Win
	}
	if you == Scissors && them == Paper {
		return Win
	}
	return Loss
}

func findYourMove(them Move, outcome Outcome) Move {
	moves := []Move{Rock, Paper, Scissors}
	for _, move := range moves {
		if calculateOutcome(move, them) == outcome {
			return move
		}
	}
	return Unknown
}

type Round struct {
	You     Move
	Them    Move
	Outcome Outcome
	Score   int
}

func NewRound(you, them Move, outcome Outcome) Round {
	return Round{
		You:   you,
		Them:  them,
		Score: int(outcome) + int(you),
	}
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileLines := strings.Split(string(fileContents), "\n")

	rounds := make([]Round, 0)
	for _, line := range fileLines {
		roundData := strings.Split(line, " ")
		them, outcome := getMove(roundData[0]), getOutcome(roundData[1])
		you := findYourMove(them, outcome)
		rounds = append(rounds, NewRound(you, them, outcome))
	}

	totalScore := 0
	for _, round := range rounds {
		totalScore += round.Score
	}

	fmt.Println(totalScore)
}
