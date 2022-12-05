package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Section int

type Assignment struct {
	Start Section
	End   Section
}

func (a Assignment) Overlaps(b Assignment) bool {
	return a.Start >= b.Start && a.End <= b.End
}

func NewAssignment(start, end string) *Assignment {
	startSection, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		panic(err)
	}
	endSection, err := strconv.ParseInt(end, 10, 64)
	if err != nil {
		panic(err)
	}
	return &Assignment{
		Start: Section(startSection),
		End:   Section(endSection),
	}
}

type AssignmentPair struct {
	Assignment1 Assignment
	Assignment2 Assignment
	Overlaps    bool
}

func NewAssignmentPair(assignment1, assignment2 Assignment) *AssignmentPair {
	return &AssignmentPair{
		Assignment1: assignment1,
		Assignment2: assignment2,
		Overlaps:    assignment1.Overlaps(assignment2) || assignment2.Overlaps(assignment1),
	}
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	assignmentPairs := make([]*AssignmentPair, len(lines))
	for i, line := range lines {
		assignments := strings.Split(line, ",")

		assignmentRange1 := strings.Split(assignments[0], "-")
		assignment1 := NewAssignment(assignmentRange1[0], assignmentRange1[1])

		assignmentRange2 := strings.Split(assignments[1], "-")
		assignment2 := NewAssignment(assignmentRange2[0], assignmentRange2[1])

		assignmentPairs[i] = NewAssignmentPair(*assignment1, *assignment2)
	}

	numberOverlapping := 0
	for _, assignmentPair := range assignmentPairs {
		if assignmentPair.Overlaps {
			numberOverlapping++
		}
	}

	fmt.Println(numberOverlapping)
}
