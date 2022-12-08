package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	trees := make([][]int, len(lines))
	for i, line := range lines {
		trees[i] = make([]int, len(line))
		for j, tree := range line {
			treeHeight, err := strconv.Atoi(string(tree))
			if err != nil {
				panic(err)
			}
			trees[i][j] = treeHeight
		}
	}

	// visible from left
	visibleLeftGrid := make([][]bool, len(trees))
	for i, row := range trees {
		visibleLeftGrid[i] = make([]bool, len(row))
		maxHeightSeen := -1
		for j, treeHeight := range row {
			if treeHeight > maxHeightSeen {
				visibleLeftGrid[i][j] = true
				maxHeightSeen = treeHeight
			}
		}
	}

	// visible from right
	visibleRightGrid := make([][]bool, len(trees))
	for i, row := range trees {
		visibleRightGrid[i] = make([]bool, len(row))
		maxHeightSeen := -1
		for j := len(row) - 1; j >= 0; j-- {
			treeHeight := row[j]
			if treeHeight > maxHeightSeen {
				visibleRightGrid[i][j] = true
				maxHeightSeen = treeHeight
			}
		}
	}

	// visible from top
	visibleTopGrid := make([][]bool, len(trees))
	for i := 0; i < len(trees); i++ {
		visibleTopGrid[i] = make([]bool, len(trees[i]))
	}

	for j := 0; j < len(trees[0]); j++ {
		maxHeightSeen := -1
		for i := 0; i < len(trees); i++ {
			treeHeight := trees[i][j]
			if treeHeight > maxHeightSeen {
				visibleTopGrid[i][j] = true
				maxHeightSeen = treeHeight
			}
		}
	}

	// visible from bottom
	visibleBottomGrid := make([][]bool, len(trees))
	for i := 0; i < len(trees); i++ {
		visibleBottomGrid[i] = make([]bool, len(trees[i]))
	}

	for j := 0; j < len(trees[0]); j++ {
		maxHeightSeen := -1
		for i := len(trees) - 1; i >= 0; i-- {
			treeHeight := trees[i][j]
			if treeHeight > maxHeightSeen {
				visibleBottomGrid[i][j] = true
				maxHeightSeen = treeHeight
			}
		}
	}

	numVisible := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			if visibleLeftGrid[i][j] || visibleRightGrid[i][j] || visibleTopGrid[i][j] || visibleBottomGrid[i][j] {
				numVisible++
			}
		}
	}

	fmt.Println(numVisible)
}
