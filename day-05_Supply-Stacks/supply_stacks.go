package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseStacks(data string, max int) map[int]string {
	tmpStacks := []string(nil)
	for _, line := range strings.Split(data, "\n") {
		if line[1] == '1' {
			break
		}
		tmpLine := ""
		for i := 0; i < max*4; i += 4 {
			if i < len(line) && line[i] == '[' {
				tmpLine += string(line[i+1])
			} else {
				tmpLine += "-"
			}
		}
		tmpStacks = append(tmpStacks, tmpLine)
	}
	reverseStacks := []string(nil)
	for i := len(tmpStacks) - 1; i >= 0; i-- {
		reverseStacks = append(reverseStacks, tmpStacks[i])
	}
	stacks := map[int]string{}
	for i := 0; i < max; i++ {
		stack := ""
		for _, reverseStack := range reverseStacks {
			if reverseStack[i] == '-' {
				break
			}
			stack += string(reverseStack[i])
		}
		stacks[i+1] = stack
	}
	return stacks
}

func performMoves(data string, stacks map[int]string) map[int]string {
	for _, instruction := range strings.Split(data, "\n") {
		if !strings.HasPrefix(instruction, "move") {
			continue
		}
		quantity := 0
		from := 0
		to := 0
		_, err := fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &from, &to)
		if err != nil {
			log.Fatalf("error scanning instructions line %s: %v", instruction, err)
		}
		toMove := stacks[from][len(stacks[from])-quantity:]
		for i := quantity - 1; i >= 0; i-- {
			stacks[to] += string(toMove[i])
		}
		stacks[from] = strings.TrimRight(stacks[from], toMove)
	}
	return stacks
}

func main() {
	//buffer, err := os.ReadFile("stacks_test.txt")
	buffer, err := os.ReadFile("stacks_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	stacksInit := parseStacks(string(buffer), 9)
	stacksFinal := performMoves(string(buffer), stacksInit)

	finalStr := ""
	for _, stack := range stacksFinal {
		finalStr += string(stack[len(stack)-1])
	}
	fmt.Printf("Final state of the stacks: %s", finalStr)
}
