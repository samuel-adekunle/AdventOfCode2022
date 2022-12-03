package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Item struct {
	Type     rune
	Priority int
}

func NewItem(item rune) *Item {
	var priority int

	if unicode.IsUpper(item) {
		priority = int(item) - 'A' + 1 + 26
	} else {
		priority = int(item) - 'a' + 1
	}

	if priority < 1 || priority > 52 {
		log.Panicf("invalid priority %d from item %v", priority, item)
	}

	return &Item{
		Type:     item,
		Priority: priority,
	}
}

type Compartment []*Item

func NewCompartment(items string) Compartment {
	var compartment Compartment
	for _, item := range items {
		compartment = append(compartment, NewItem(item))
	}
	return compartment
}

type Rucksack [2]Compartment

func (r *Rucksack) findSharedItem() *Item {
	for _, item := range r[0] {
		for _, item2 := range r[1] {
			if item.Type == item2.Type {
				return item
			}
		}
	}
	return nil
}

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileLines := strings.Split(string(fileBytes), "\n")

	rucksacks := make([]Rucksack, len(fileLines))

	for i, line := range fileLines {
		n := len(line)
		compartment1, compartment2 := line[:n/2], line[n/2:]
		if len(compartment1) != len(compartment2) {
			panic("compartment lengths are not equal")
		}
		rucksacks[i] = Rucksack{NewCompartment(compartment1), NewCompartment(compartment2)}
	}

	totalSharedItemsPriorities := 0

	for _, r := range rucksacks {
		sharedItem := r.findSharedItem()
		fmt.Println(sharedItem)
		totalSharedItemsPriorities += sharedItem.Priority
	}

	fmt.Println(totalSharedItemsPriorities)
}
