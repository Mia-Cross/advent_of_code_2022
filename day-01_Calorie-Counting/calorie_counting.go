package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countCaloriesByElf(data string) []int {
	sums := []int(nil)
	sum := 0
	for _, numberStr := range strings.Split(data, "\n") {
		if numberStr != "" {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}
	return sums
}

func findMostCarryingElves(sums []int) []int {
	topThree := []int{0, 0, 0}
	for _, sum := range sums {
		if sum > topThree[0] {
			topThree[2] = topThree[1]
			topThree[1] = topThree[0]
			topThree[0] = sum
		} else if sum > topThree[1] {
			topThree[2] = topThree[1]
			topThree[1] = sum
		} else if sum > topThree[2] {
			topThree[2] = sum
		}
	}
	return topThree[:3]
}

func main() {
	buffer, err := os.ReadFile("calorie_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	calorieSums := countCaloriesByElf(string(buffer))
	topThree := findMostCarryingElves(calorieSums)

	fmt.Printf("Total calories carried by the first elf: %d\n", topThree[0])
	fmt.Printf("Total calories carried by the top 3 elves: %d\n", topThree[0]+topThree[1]+topThree[2])
}

//Total calories carried by the first elf: 71124
//Total calories carried by the top 3 elves: 204639
