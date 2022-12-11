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
	output := [6][41]byte{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			output[i][j] = '.'
		}
		output[i][40] = '\n'
	}

	spriteIndex := -1
	previousValue := 1
	lineIndex := 0
	expected := "##..##..##..##..##..##..##..##..##..##.."

	for i, value := range register {
		//for i := 0; i < len(register); i++ {
		//	value := register[i]
		if spriteIndex == 1 && previousValue == value || value < 0 {
			continue
		}
		if previousValue != value {
			spriteIndex = -1
		}
		if value%2 == 0 {
			indexToWrite := value + spriteIndex + 1
			output[lineIndex][indexToWrite] = '#'
			if output[lineIndex][indexToWrite] != expected[indexToWrite] {
				fmt.Printf("hop hop hop ! @ index %d du registre", i)
			}
		} else {
			indexToWrite := value + spriteIndex
			output[lineIndex][indexToWrite] = '#'
			if output[lineIndex][indexToWrite] != expected[indexToWrite] {
				fmt.Printf("hop hop hop ! @ index %d du registre", i)
			}
		}
		spriteIndex++
		if (i+1)%40 == 0 {
			lineIndex++
		}
		previousValue = value
	}

	outputStr := [6]string{}
	for i := range output {
		str := ""
		for j := range output[i] {
			str += string(output[i][j])
		}
		outputStr[i] = str
	}
	return outputStr
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
		fmt.Printf("%s\n", line)
	}
}
