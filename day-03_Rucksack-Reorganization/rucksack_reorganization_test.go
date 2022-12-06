package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateScoreForItems(t *testing.T) {
	buffer, err := os.ReadFile("items_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	assert.Equal(t, 157, calculatePrioritiesForItems(string(buffer)))
}

func TestCalculateScoreForBadges(t *testing.T) {
	buffer, err := os.ReadFile("items_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	assert.Equal(t, 70, calculatePrioritiesForBadges(string(buffer)))
}
