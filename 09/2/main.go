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

func keepAdjacent(head, tail *Position) {
	xDelta, yDelta := head.X-tail.X, head.Y-tail.Y
	if abs(xDelta) >= 2 || abs(yDelta) >= 2 {
		// move tail closer to head

		// head directly to the left
		if xDelta >= 2 && yDelta == 0 {
			tail.X += xDelta - 1
		}

		// head directly to the right
		if xDelta <= -2 && yDelta == 0 {
			tail.X -= xDelta + 1
		}

		// head directly above
		if xDelta == 0 && yDelta >= 2 {
			tail.Y += yDelta - 1
		}

		// head directly below
		if xDelta == 0 && yDelta <= -2 {
			tail.Y -= yDelta + 1
		}

		// head diagonally up and to the left
		if xDelta >= 2 && yDelta >= 1 && abs(xDelta) > abs(yDelta) {
			tail.X += xDelta - 1
			tail.Y += yDelta
		}
		if xDelta >= 1 && yDelta >= 2 && abs(xDelta) < abs(yDelta) {
			tail.X += xDelta
			tail.Y += yDelta - 1
		}
		if xDelta >= 2 && yDelta >= 2 && abs(xDelta) == abs(yDelta) {
			tail.X += xDelta - 1
			tail.Y += yDelta - 1
		}

		// head diagonally up and to the right
		if xDelta <= -2 && yDelta >= 1 && abs(xDelta) > yDelta {
			tail.X -= xDelta + 1
			tail.Y += yDelta
		}
		if xDelta <= -1 && yDelta >= 2 && abs(xDelta) < yDelta {
			tail.X -= xDelta
			tail.Y += yDelta - 1
		}
		if xDelta <= -2 && yDelta >= 2 && abs(xDelta) == abs(yDelta) {
			tail.X -= xDelta + 1
			tail.Y += yDelta - 1
		}

		// head diagonally down and to the left
		if xDelta >= 2 && yDelta <= -1 && abs(xDelta) > abs(yDelta) {
			tail.X += xDelta - 1
			tail.Y += yDelta
		}
		if xDelta >= 1 && yDelta <= -2 && abs(xDelta) < abs(yDelta) {
			tail.X += xDelta
			tail.Y += yDelta + 1
		}
		if xDelta >= 2 && yDelta <= -2 && abs(xDelta) == abs(yDelta) {
			tail.X += xDelta - 1
			tail.Y += yDelta + 1
		}

		// head diagonally down and to the right
		if xDelta <= -2 && yDelta <= -1 && abs(xDelta) > abs(yDelta) {
			tail.X -= xDelta + 1
			tail.Y += yDelta
		}
		if xDelta == -1 && yDelta == -2 && abs(xDelta) < abs(yDelta) {
			tail.X -= xDelta
			tail.Y += yDelta + 1
		}
		if xDelta <= -2 && yDelta <= -2 && abs(xDelta) == abs(yDelta) {
			tail.X -= xDelta + 1
			tail.Y += yDelta + 1
		}
	}
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.ex.txt")
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

	rope := make([]*Position, 10)
	for i := 0; i < len(rope); i++ {
		rope[i] = &Position{0, 0}
	}

	head := func() *Position {
		return rope[0]
	}

	tail := func() *Position {
		return rope[len(rope)-1]
	}

	// track positions visited by the tail
	visited := make(map[Position]bool)
	visited[*tail()] = true

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			switch move.Direction {
			case Up:
				head().Y++
			case Down:
				head().Y--
			case Left:
				head().X--
			case Right:
				head().X++
			}

			for i := 0; i < len(rope)-1; i++ {
				_head, _tail := rope[i], rope[i+1]
				keepAdjacent(_head, _tail)
			}
			visited[*tail()] = true
		}
	}

	fmt.Println(len(visited))
}
