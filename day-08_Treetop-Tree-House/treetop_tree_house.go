package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func isVisible(grid []string, x, y int) int {
	height := grid[y][x]
	hidden := ""
	for i := x + 1; i < len(grid[y]); i++ {
		if grid[y][i] >= height {
			hidden += "X"
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if grid[y][i] >= height {
			hidden += "X"
			break
		}
	}
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= height {
			hidden += "X"
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] >= height {
			hidden += "X"
			break
		}
	}
	if len(hidden) == 4 {
		return 0
	}
	return 1
}

func calculateScenicScore(grid []string, x, y int) int {
	height := grid[y][x]
	scores := []int{0, 0, 0, 0}
	for i := x + 1; i < len(grid[y]); i++ {
		scores[0]++
		if grid[y][i] >= height {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		scores[1]++
		if grid[y][i] >= height {
			break
		}
	}
	for i := y + 1; i < len(grid); i++ {
		scores[2]++
		if grid[i][x] >= height {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		scores[3]++
		if grid[i][x] >= height {
			break
		}
	}
	return scores[0] * scores[1] * scores[2] * scores[3]
}

func countVisibleTrees(data string) (int, int) {
	grid := strings.Split(data, "\n")
	bestScore := 0
	count := 0
	for y := range grid {
		for x := range grid[y] {
			count += isVisible(grid, x, y)
			score := calculateScenicScore(grid, x, y)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return count, bestScore
}

func main() {
	buffer, err := os.ReadFile("trees_grid.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	visibleTrees, bestScenicScore := countVisibleTrees(string(buffer))

	fmt.Printf("Number of trees visible from outside the grid: %d\n", visibleTrees)
	fmt.Printf("Best scenic score: %d\n", bestScenicScore)
}
