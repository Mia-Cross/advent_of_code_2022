package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swapValues(limits []int) []int {
	tmp := limits[0]
	limits[0] = limits[2]
	limits[2] = tmp
	tmp = limits[1]
	limits[1] = limits[3]
	limits[3] = tmp
	return limits
}

func countOverlappingSections(limits []int) int {
	// Placing the section with the lowest minimum first
	if limits[0] > limits[2] {
		limits = swapValues(limits)
	}
	if limits[1] >= limits[2] {
		return 1
	}
	return 0
}

func countFullyContainedSections(limits []int) int {
	// Placing the section with the lowest range first
	if limits[1]-limits[0] > limits[3]-limits[2] {
		limits = swapValues(limits)
	}
	if limits[0] >= limits[2] && limits[0] <= limits[3] &&
		limits[1] >= limits[2] && limits[1] <= limits[3] {
		return 1
	}
	return 0
}

func parseSectionsLimits(data string) [][]int {
	limitsArray := [][]int(nil)
	for _, pair := range strings.Split(data, "\n") {
		sections := strings.Replace(pair, ",", "-", -1)
		limits := []int(nil)
		for _, limit := range strings.Split(sections, "-") {
			nb, err := strconv.Atoi(limit)
			if err != nil {
				log.Fatalf("could not convert [%s] to number", limit)
			}
			limits = append(limits, nb)
		}
		limitsArray = append(limitsArray, limits)
	}
	return limitsArray
}

func main() {
	buffer, err := os.ReadFile("sections_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	fullyContained := 0
	overlapping := 0
	limitsArray := parseSectionsLimits(string(buffer))
	for _, limits := range limitsArray {
		fullyContained += countFullyContainedSections(limits)
		overlapping += countOverlappingSections(limits)
	}

	fmt.Printf("Number of fully contained sections: %d\n", fullyContained)
	fmt.Printf("Number of overlaping sections: %d\n", overlapping)
}
