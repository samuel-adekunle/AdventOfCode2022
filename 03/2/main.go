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

type Rucksack Compartment

func (r *Rucksack) findSharedItems(other *Rucksack) Rucksack {
	seenItems := map[rune]*Item{}
	for _, item := range *r {
		for _, otherItem := range *other {
			if item.Type == otherItem.Type {
				seenItems[item.Type] = item
			}
		}
	}

	var sharedItems Rucksack
	for _, item := range seenItems {
		sharedItems = append(sharedItems, item)
	}

	return sharedItems
}

type ElfGroup [3]Rucksack

func (e *ElfGroup) findGroupBadge() *Item {
	sharedItems := e[0].findSharedItems(&e[1])
	sharedItems = sharedItems.findSharedItems(&e[2])
	if len(sharedItems) != 1 {
		log.Panicf("expected 1 shared item, got %d", len(sharedItems))
	}
	return sharedItems[0]
}

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileLines := strings.Split(string(fileBytes), "\n")

	rucksacks := make([]Rucksack, len(fileLines))

	for i, line := range fileLines {
		rucksacks[i] = Rucksack(NewCompartment(line))
	}

	elfGroups := make([]ElfGroup, len(rucksacks)/3)

	for i := 0; i < len(rucksacks); i += 3 {
		elfGroups[i/3] = ElfGroup{rucksacks[i], rucksacks[i+1], rucksacks[i+2]}
	}

	totalGroupBadgePriority := 0
	for _, elfGroup := range elfGroups {
		totalGroupBadgePriority += elfGroup.findGroupBadge().Priority
	}

	fmt.Println(totalGroupBadgePriority)
}
