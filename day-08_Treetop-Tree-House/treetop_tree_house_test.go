package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVisibleTrees(t *testing.T) {
	buffer, err := os.ReadFile("trees_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	visibleTrees, _ := countVisibleTrees(string(buffer))

	assert.Equal(t, 21, visibleTrees)
}

func TestFindBestTree(t *testing.T) {
	buffer, err := os.ReadFile("trees_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	_, scenicScore := countVisibleTrees(string(buffer))

	assert.Equal(t, 8, scenicScore)
}
