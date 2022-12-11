package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func addIfNew(visited [][2]int, toAdd [2]int) [][2]int {
	for _, pair := range visited {
		if pair[0] == toAdd[0] && pair[1] == toAdd[1] {
			return visited
		}
	}
	return append(visited, toAdd)
}

func moveTail(head, tail [2]int) [2]int {
	if head[0] > tail[0] {
		if head[0] > 0 {

		} else {

		}
	}
}

func moveRope(direction byte, nb int, head, tail [2]int) ([2]int, [2]int) {
	for i := 0; i < nb; i++ {
		switch direction {
		case 'R':
			head[0]++
		case 'L':
			head[0]--
		case 'U':
			head[1]++
		case 'D':
			head[1]--
		}
		tail = moveTail(head, tail)
	}
	return head, tail
}

func countVisitedCells(data string) int {
	visited := [][2]int(nil)
	head := [2]int{0, 0}
	tail := [2]int{0, 0}
	for _, move := range strings.Split(data, "\n") {
		nb, err := strconv.Atoi(move[2:])
		if err != nil {
			log.Fatalf("could not parse [%s]: %v", move, err)
		}
		head, tail = moveRope(move[0], nb, head, tail)
		addIfNew(visited, tail)
	}
	return len(visited)
}

func main() {
	buffer, err := os.ReadFile("rope_moves.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	visited := countVisitedCells(string(buffer))

	fmt.Printf("Number of cells visited by the tail of the rope: %d", visited)
}
