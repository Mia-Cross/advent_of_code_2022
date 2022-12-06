package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountFullyContainedSections(t *testing.T) {
	buffer, err := os.ReadFile("sections_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	fullyContained := 0
	limitsArray := parseSectionsLimits(string(buffer))
	for _, limits := range limitsArray {
		fullyContained += countFullyContainedSections(limits)
	}
	assert.Equal(t, 2, fullyContained)
}

func TestCountOverlappingSections(t *testing.T) {
	buffer, err := os.ReadFile("sections_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	overlapping := 0
	limitsArray := parseSectionsLimits(string(buffer))
	for _, limits := range limitsArray {
		overlapping += countOverlappingSections(limits)
	}
	assert.Equal(t, 6, overlapping)
}
