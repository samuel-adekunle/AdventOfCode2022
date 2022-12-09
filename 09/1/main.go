package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

type Direction rune

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

type Move struct {
	Direction Direction
	Steps     int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	moves := make([]Move, len(lines))

	for i, line := range lines {
		move := strings.Split(line, " ")
		direction, steps := Direction(move[0][0]), move[1]
		numSteps, err := strconv.Atoi(string(steps))
		if err != nil {
			panic(err)
		}
		moves[i] = Move{direction, numSteps}
	}

	head, tail := Position{0, 0}, Position{0, 0}
	visited := make(map[Position]bool)
	visited[tail] = true

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			switch move.Direction {
			case Up:
				head.Y++
			case Down:
				head.Y--
			case Left:
				head.X--
			case Right:
				head.X++
			}

			xDelta, yDelta := head.X-tail.X, head.Y-tail.Y
			if abs(xDelta) == 2 || abs(yDelta) == 2 {
				// move tail closer to head

				// head directly to the left
				if xDelta == 2 && yDelta == 0 {
					tail.X++
				}

				// head directly to the right
				if xDelta == -2 && yDelta == 0 {
					tail.X--
				}

				// head directly above
				if xDelta == 0 && yDelta == 2 {
					tail.Y++
				}

				// head directly below
				if xDelta == 0 && yDelta == -2 {
					tail.Y--
				}

				// head diagonally up and to the left
				if xDelta == 2 && yDelta == 1 || xDelta == 1 && yDelta == 2 {
					tail.X++
					tail.Y++
				}

				// head diagonally up and to the right
				if xDelta == -2 && yDelta == 1 || xDelta == -1 && yDelta == 2 {
					tail.X--
					tail.Y++
				}

				// head diagonally down and to the left
				if xDelta == 2 && yDelta == -1 || xDelta == 1 && yDelta == -2 {
					tail.X++
					tail.Y--
				}

				// head diagonally down and to the right
				if xDelta == -2 && yDelta == -1 || xDelta == -1 && yDelta == -2 {
					tail.X--
					tail.Y--
				}

				visited[tail] = true
			}

		}
	}

	fmt.Println(len(visited))
}
