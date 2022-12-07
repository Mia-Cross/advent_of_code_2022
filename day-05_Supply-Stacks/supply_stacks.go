package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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

func performMoves(data string, stacks map[int]string) map[int]string {
	//printStacks(stacks)
	quantity := 0
	from := 0
	to := 0
	for _, instruction := range strings.Split(data, "\n") {
		//fmt.Printf("\n__________________\nEXECUTING -> %s\n\n", instruction)
		if !strings.HasPrefix(instruction, "move") {
			continue
		}
		_, err := fmt.Sscanf(instruction, "move %d from %d to %d", &quantity, &from, &to)
		if err != nil {
			log.Fatalf("error scanning instructions line %s: %v", instruction, err)
		}
		//if len(stacks[from])-quantity < 0 {
		//	continue
		//}
		oldStacks := stacks
		toMove := stacks[from][len(stacks[from])-quantity:]
		for i := quantity - 1; i >= 0; i-- {
			stacks[to] += string(toMove[i])
		}
		stacks[from] = strings.TrimRight(stacks[from], toMove)
		if !checkStacks(stacks, oldStacks, quantity, from, to) {
			log.Fatalf("NOPE !! move %d from %d to %d fucked up", quantity, from, to)
		}
		//printStacks(stacks)
	}
	return stacks
}

func main() {
	buffer, err := os.ReadFile("stacks_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	stacksInit := parseStacks(string(buffer), 9)
	stacksFinal := performMoves(string(buffer), stacksInit)
	//printStacks(stacksFinal)

	finalStr := ""
	for i := 1; i < len(stacksFinal)+1; i++ {
		//if stacksFinal[i] == "" {
		//	continue
		//}
		finalStr += string(stacksFinal[i][len(stacksFinal[i])-1])
	}
	fmt.Printf("Final state of the stacks: %s", finalStr)
}
