package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Directory struct {
	name           string
	filesSize      int
	subDirectories []Directory
	//parentDirectory Directory
}

func calculateTotalFilesSize(data string) int {
	sum := 0
	fileTree := Directory{}
	for i, instruction := range strings.Split(data, "\n") {
		switch instruction {
		case "$ ls":
			build
		}
		if instruction == "$ ls" {
			fileTree
		}
	}
	return sum
}

func main() {
	buffer, err := os.ReadFile("commands_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	size := calculateTotalFilesSize(string(buffer))

	fmt.Printf("Total size of directories over 100.000: %d", size)
}
