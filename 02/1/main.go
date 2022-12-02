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
	if move == "A" || move == "X" {
		return Rock
	}
	if move == "B" || move == "Y" {
		return Paper
	}
	if move == "C" || move == "Z" {
		return Scissors
	}
	return Unknown
}

func calculateYourScore(you, them Move) int {
	if you == them {
		return 3
	}
	if you == Rock && them == Scissors {
		return 6
	}
	if you == Paper && them == Rock {
		return 6
	}
	if you == Scissors && them == Paper {
		return 6
	}
	return 0
}

type Round struct {
	You       Move
	Them      Move
	YourScore int
}

func NewRound(you, them Move) Round {
	return Round{
		You:       you,
		Them:      them,
		YourScore: calculateYourScore(you, them) + int(you),
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
		moves := strings.Split(line, " ")
		them, you := getMove(moves[0]), getMove(moves[1])
		rounds = append(rounds, NewRound(you, them))
	}

	totalScore := 0
	for _, round := range rounds {
		totalScore += round.YourScore
	}

	fmt.Println(totalScore)
}
