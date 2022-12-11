package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalFilesSize(t *testing.T) {
	buffer, err := os.ReadFile("commands_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	size := calculateTotalFilesSize(string(buffer))
	assert.Equal(t, 95437, size)
}
