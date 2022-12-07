package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransmissionZero(t *testing.T) {
	marker := findFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)
	message := findFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
	assert.Equal(t, 7, marker)
	assert.Equal(t, 19, message)
}

func TestTransmissionOne(t *testing.T) {
	marker := findFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	message := findFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14)
	assert.Equal(t, 5, marker)
	assert.Equal(t, 23, message)
}

func TestTransmissionTwo(t *testing.T) {
	marker := findFirstMarker("nppdvjthqldpwncqszvftbrmjlhg", 4)
	message := findFirstMarker("nppdvjthqldpwncqszvftbrmjlhg", 14)
	assert.Equal(t, 6, marker)
	assert.Equal(t, 23, message)
}

func TestTransmissionThree(t *testing.T) {
	marker := findFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	message := findFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14)
	assert.Equal(t, 10, marker)
	assert.Equal(t, 29, message)
}

func TestTransmissionFour(t *testing.T) {
	marker := findFirstMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
	message := findFirstMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14)
	assert.Equal(t, 11, marker)
	assert.Equal(t, 26, message)
}
