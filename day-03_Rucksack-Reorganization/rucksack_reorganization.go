package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func calculatePrioritiesForItems(data string) int {
	sum := 0
	for _, items := range strings.Split(data, "\n") {
		var duplicateItem rune
		for _, item := range items[:len(items)/2] {
			if strings.Contains(items[len(items)/2:], string(item)) {
				duplicateItem = item
				break
			}
		}
		if unicode.IsLower(duplicateItem) {
			sum += int(duplicateItem - 'a' + 1)
		} else {
			sum += int(duplicateItem - 'A' + 27)
		}
	}
	return sum
}

func calculatePrioritiesForBadges(data string) int {
	sum := 0
	groups := map[int][]string{}
	groupIndex := 0
	var duplicateItem rune

	for i, items := range strings.Split(data, "\n") {
		groups[groupIndex] = append(groups[groupIndex], items)
		if (i+1)%3 == 0 {
			groupIndex++
		}
	}

	for _, group := range groups {
		commonChars := ""
		for _, item := range group[0] {
			if strings.Contains(group[1], string(item)) {
				commonChars += string(item)
			}
		}
		for _, commonChar := range commonChars {
			if strings.Contains(group[2], string(commonChar)) {
				duplicateItem = commonChar
				break
			}
		}

		if unicode.IsLower(duplicateItem) {
			sum += int(duplicateItem - 'a' + 1)
		} else {
			sum += int(duplicateItem - 'A' + 27)
		}
	}

	return sum
}

func main() {
	buffer, err := os.ReadFile("items_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	fmt.Printf("Sum of the priority scores of the items: %d\n", calculatePrioritiesForItems(string(buffer)))
	fmt.Printf("Sum of the priority scores of the badges: %d\n", calculatePrioritiesForBadges(string(buffer)))
}
