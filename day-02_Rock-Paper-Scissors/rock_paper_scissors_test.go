package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateScoreNewRules(t *testing.T) {
	buffer, err := os.ReadFile("moves_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	assert.Equal(t, 12, calculateScoreNewRules(string(buffer)))
}
