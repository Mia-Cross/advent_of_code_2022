package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVisitedCells(t *testing.T) {
	buffer, err := os.ReadFile("rope_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	visited := countVisitedCells(string(buffer))
	assert.Equal(t, 13, visited)
}
