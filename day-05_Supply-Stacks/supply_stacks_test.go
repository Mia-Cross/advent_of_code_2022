package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const numberOfStacksTest = 3

func TestSupplyStacks9000(t *testing.T) {
	buffer, err := os.ReadFile("stacks_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	stacksInit := parseStacks(string(buffer), numberOfStacksTest)
	stacksFinal := performMoves(string(buffer), stacksInit, 9000)

	assert.Equal(t, "CMZ", stacksFinal)
}

func TestSupplyStacks9001(t *testing.T) {
	buffer, err := os.ReadFile("stacks_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	stacksInit := parseStacks(string(buffer), numberOfStacksTest)
	stacksFinal := performMoves(string(buffer), stacksInit, 9001)

	assert.Equal(t, "MCD", stacksFinal)
}
