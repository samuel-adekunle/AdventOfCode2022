package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Crate rune

type Stack []Crate

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(c Crate) {
	*s = append(*s, c)
}

// reverse the stack
func (s *Stack) Reverse() {
	for i := 0; i < len(*s)/2; i++ {
		j := len(*s) - i - 1
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func (s *Stack) Pop() Crate {
	c := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return c
}

func (s *Stack) Peek() Crate {
	return (*s)[len(*s)-1]
}

type MovementInstruction struct {
	From  int
	To    int
	Count int
}

func NewMovementInstruction(instructionLine string) *MovementInstruction {
	details := strings.Split(instructionLine, " ")
	count, err := strconv.ParseInt(details[1], 10, 64)
	if err != nil {
		panic(err)
	}
	from, err := strconv.ParseInt(details[3], 10, 64)
	if err != nil {
		panic(err)
	}
	to, err := strconv.ParseInt(details[5], 10, 64)
	if err != nil {
		panic(err)
	}
	return &MovementInstruction{
		From:  int(from - 1),
		To:    int(to - 1),
		Count: int(count),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	readMode := 0
	stacks := make(map[int]*Stack, 0)
	movementInstructions := make([]*MovementInstruction, 0)
	for _, line := range lines {
		if line == "" {
			readMode++

			// reverse stacks
			for i := 0; i < len(stacks); i++ {
				stacks[i].Reverse()
			}

			continue
		}

		switch readMode {
		case 0:
			for i := 0; i < len(line); i += 4 {
				if line[i] != '[' {
					continue
				}
				stack := i / 4
				crate := Crate(line[i+1])
				if _, ok := stacks[stack]; !ok {
					stacks[stack] = NewStack()
				}
				stacks[stack].Push(crate)
			}
		case 1:
			movementInstructions = append(movementInstructions, NewMovementInstruction(line))
		}
	}

	for _, instruction := range movementInstructions {
		for i := 0; i < instruction.Count; i++ {
			stacks[instruction.To].Push(stacks[instruction.From].Pop())
		}
	}

	cratesAtTop := make([]Crate, len(stacks))
	for i := 0; i < len(stacks); i++ {
		cratesAtTop[i] = stacks[i].Peek()
	}

	fmt.Println(string(cratesAtTop))
}
