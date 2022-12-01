package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Food struct {
	calories int
}

type Elf struct {
	food          []Food
	totalCalories int
}

func NewElf() *Elf {
	elf := new(Elf)
	elf.food = make([]Food, 0)
	elf.totalCalories = 0
	return elf
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	elves, curElf := make([]*Elf, 0), NewElf()
	for _, line := range strings.Split(string(f), "\n") {
		if line == "" {
			elves = append(elves, curElf)
			curElf = NewElf()
			continue
		}

		calories, err := strconv.ParseInt(line, 10, 0)
		if err != nil {
			panic(err)
		}

		curElf.food = append(curElf.food, Food{calories: int(calories)})
		curElf.totalCalories += int(calories)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].totalCalories < elves[j].totalCalories
	})

	n := len(elves)

	// Part 1s
	fmt.Println(elves[n-1].totalCalories)

	// Part 2
	totalCalories := 0
	for _, elf := range elves[n-3:] {
		totalCalories += elf.totalCalories
	}

	fmt.Println(totalCalories)
}
