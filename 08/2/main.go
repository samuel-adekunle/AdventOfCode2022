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

	// number visible from left
	visibleLeftGrid := make([][]int, len(trees))
	for i, row := range trees {
		visibleLeftGrid[i] = make([]int, len(row))
		for j, treeHeight := range row {
			for k := j - 1; k >= 0; k-- {
				visibleLeftGrid[i][j]++
				if treeHeight <= row[k] {
					break
				}
			}
		}
	}

	// number visible from right
	visibleRightGrid := make([][]int, len(trees))
	for i, row := range trees {
		visibleRightGrid[i] = make([]int, len(row))
		for j := len(row) - 1; j >= 0; j-- {
			treeHeight := row[j]
			for k := j + 1; k < len(row); k++ {
				visibleRightGrid[i][j]++
				if treeHeight <= row[k] {
					break
				}
			}
		}
	}

	// number visible from top
	visibleTopGrid := make([][]int, len(trees))
	for i := 0; i < len(trees); i++ {
		visibleTopGrid[i] = make([]int, len(trees[i]))
	}

	for j := 0; j < len(trees[0]); j++ {
		for i := 0; i < len(trees); i++ {
			treeHeight := trees[i][j]
			for k := i - 1; k >= 0; k-- {
				visibleTopGrid[i][j]++
				if treeHeight <= trees[k][j] {
					break
				}
			}
		}
	}

	// number visible from bottom
	visibleBottomGrid := make([][]int, len(trees))
	for i := 0; i < len(trees); i++ {
		visibleBottomGrid[i] = make([]int, len(trees[i]))
	}

	for j := 0; j < len(trees[0]); j++ {
		for i := len(trees) - 1; i >= 0; i-- {
			treeHeight := trees[i][j]
			for k := i + 1; k < len(trees); k++ {
				visibleBottomGrid[i][j]++
				if treeHeight <= trees[k][j] {
					break
				}
			}
		}
	}

	// scenic score = number visible from left * number visible from right * number visible from top * number visible from bottom

	maxScenicScore := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			scenicScore := visibleLeftGrid[i][j] * visibleRightGrid[i][j] * visibleTopGrid[i][j] * visibleBottomGrid[i][j]
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println(maxScenicScore)
}
