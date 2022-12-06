package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupplyStacks(t *testing.T) {
	buffer, err := os.ReadFile("stacks_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	stacksInit := parseStacks(string(buffer), 3)
	stacksFinal := performMoves(string(buffer), stacksInit)

	finalStr := ""
	for _, stack := range stacksFinal {
		finalStr += string(stack[len(stack)-1])
	}
	assert.Equal(t, "CMZ", finalStr)
}
