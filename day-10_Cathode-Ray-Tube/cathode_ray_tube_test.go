package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateKeySignalStrength(t *testing.T) {
	buffer, err := os.ReadFile("program_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
	signalStrengths, _ := calculateKeySignalStrengths(string(buffer))
	assert.Equal(t, 13140, signalStrengths)
}

func TestDrawOutput(t *testing.T) {
	buffer, err := os.ReadFile("program_test.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	_, register := calculateKeySignalStrengths(string(buffer))
	output := drawOutput(register)

	assert.Equal(t, "##..##..##..##..##..##..##..##..##..##..\n", output[0])
	assert.Equal(t, "###...###...###...###...###...###...###.\n", output[1])
	assert.Equal(t, "####....####....####....####....####....\n", output[2])
	assert.Equal(t, "#####.....#####.....#####.....#####.....\n", output[3])
	assert.Equal(t, "######......######......######......####\n", output[4])
	assert.Equal(t, "#######.......#######.......#######.....", output[5])
}
