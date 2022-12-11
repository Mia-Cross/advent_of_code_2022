package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateKeySignalStrengths(data string) (int, []int) {
	registerValues := []int{1}
	for _, command := range strings.Split(data, "\n") {
		registerValues = append(registerValues, registerValues[len(registerValues)-1])
		if strings.HasPrefix(command, "addx") {
			x, err := strconv.Atoi(command[5:])
			if err != nil {
				log.Fatalf("could not parse [%s]: %v", command, err)
			}
			registerValues = append(registerValues, registerValues[len(registerValues)-1]+x)
		}
	}
	sumOfKeyStrengths := 0
	for i := 20; i <= 220; i += 40 {
		sumOfKeyStrengths += i * registerValues[i-1]
	}
	return sumOfKeyStrengths, registerValues
}

func drawOutput(register []int) [6]string {
	output := [6]string{}
	spritePosition := []int{0, 1, 2}
	previousValue := 1
	lineIndex := 0
	previousLineIndex := 0

	for i, value := range register {
		if previousValue != value || previousLineIndex != lineIndex {
			adjustedValue := value + (40 * lineIndex)
			spritePosition = []int{adjustedValue - 1, adjustedValue, adjustedValue + 1}
			previousLineIndex = lineIndex
		}

		printedSprite := false
		for _, index := range spritePosition {
			if i == index {
				output[lineIndex] += "#"
				printedSprite = true
			}
		}
		if printedSprite == false {
			output[lineIndex] += "."
		}

		if (i+1)%40 == 0 {
			if lineIndex == 5 {
				break
			}
			output[lineIndex] += "\n"
			previousLineIndex = lineIndex
			lineIndex++
		}
		previousValue = value
	}

	return output
}

func main() {
	buffer, err := os.ReadFile("program.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	signalStrengths, registerValues := calculateKeySignalStrengths(string(buffer))
	fmt.Printf("Sum of the 6 signal strengths: %d", signalStrengths)

	spriteOutput := drawOutput(registerValues)
	fmt.Printf("Visual output:\n")
	for _, line := range spriteOutput {
		fmt.Printf("%s", line)
	}
}
