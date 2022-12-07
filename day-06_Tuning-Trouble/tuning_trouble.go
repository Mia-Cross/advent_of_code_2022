package main

import (
	"fmt"
	"log"
	"os"
)

func everyCharacterIsUnique(substr string) bool {
	for i, charA := range substr {
		for j, charB := range substr {
			if i == j {
				continue
			}
			if charA == charB {
				return false
			}
		}
	}
	return true
}

func findFirstMarker(transmission string, markerLength int) int {
	for i := 0; i < len(transmission)-markerLength-1; i++ {
		if everyCharacterIsUnique(transmission[i : i+markerLength]) {
			return i + markerLength
		}
	}
	return 0
}

func main() {
	buffer, err := os.ReadFile("transmission.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	fmt.Printf("First marker in transmission is at index: %d\n", findFirstMarker(string(buffer), 4))
	fmt.Printf("First message in transmission is at index: %d\n", findFirstMarker(string(buffer), 14))
}
