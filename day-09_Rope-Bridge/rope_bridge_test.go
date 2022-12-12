package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVisitedCellsWithTwoKnots(t *testing.T) {
	buffer, err := os.ReadFile("rope_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	visited := countVisitedCells(string(buffer), 2)
	assert.Equal(t, 13, visited)
}

func TestCountVisitedCellsWithTenKnots(t *testing.T) {
	buffer, err := os.ReadFile("rope_10_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	visited := countVisitedCells(string(buffer), 10)
	assert.Equal(t, 36, visited)
}
