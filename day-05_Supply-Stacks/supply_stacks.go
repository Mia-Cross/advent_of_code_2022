package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const numberOfStacks = 9

func printStacks(stacks map[int]string) {
	for i := 1; i < len(stacks)+1; i++ {
		fmt.Printf("[%d] [%s]\n", i, stacks[i])
	}
}

func checkStacks(newStacks, oldStacks map[int]string, quantity, from, to int) bool {
	for i := 1; i < len(newStacks)+1; i++ {
		switch i {
		case from:
			if len(newStacks[i]) != len(oldStacks[i])-quantity {
				fmt.Printf("I expected newStack[%d] to be %d characters long : new {%s}{%s} old\n", i, len(oldStacks[i])-quantity, newStacks[i], oldStacks[i])
				return false
			}
		case to:
			if len(newStacks[i]) != len(oldStacks[i])+quantity {
				fmt.Printf("I expected newStack[%d] to be %d characters long : new {%s}{%s} old\n", i, len(oldStacks[i])+quantity, newStacks[i], oldStacks[i])
				return false
			}
		default:
			if len(newStacks[i]) != len(oldStacks[i]) {
				fmt.Printf("I expected newStack[%d] to be %d characters long : new {%s}{%s} old\n", i, len(oldStacks[i]), newStacks[i], oldStacks[i])
				return false
			}
		}
	}
	return true
}

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

func crateMover9000(stacks map[int]string, quantity, from, to int) map[int]string {
	toMove := stacks[from][len(stacks[from])-quantity:]
	for i := quantity - 1; i >= 0; i-- {
		stacks[to] += string(toMove[i])
	}
	stacks[from] = stacks[from][:len(stacks[from])-quantity]
	return stacks
}

func crateMover9001(stacks map[int]string, quantity, from, to int) map[int]string {
	toMove := stacks[from][len(stacks[from])-quantity:]
	stacks[to] += toMove
	stacks[from] = stacks[from][:len(stacks[from])-quantity]
	return stacks
}

func performMoves(data string, stacks map[int]string, craneModel int) string {
	quantity := 0
	from := 0
	to := 0

	for _, instruction := range strings.Split(data, "\n") {
		if !strings.HasPrefix(instruction, "move") {
			continue
		}
		_, err := fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &from, &to)
		if err != nil {
			log.Fatalf("error scanning instructions line %s: %v", instruction, err)
		}
		if craneModel == 9000 {
			stacks = crateMover9000(stacks, quantity, from, to)
		} else {
			stacks = crateMover9001(stacks, quantity, from, to)
		}
	}

	topCrates := ""
	for i := 1; i < len(stacks)+1; i++ {
		topCrates += string(stacks[i][len(stacks[i])-1])
	}
	return topCrates
}

func main() {
	buffer, err := os.ReadFile("stacks_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	data := string(buffer)

	stacksFinal9000 := performMoves(data, parseStacks(data, numberOfStacks), 9000)
	stacksFinal9001 := performMoves(data, parseStacks(data, numberOfStacks), 9001)

	fmt.Printf("Final state of the stacks with the CrateMover9000: %s\n", stacksFinal9000)
	fmt.Printf("Final state of the stacks with the CrateMover9001: %s\n", stacksFinal9001)
}
