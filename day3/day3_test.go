package day3

import (
	"fmt"
	"testing"
)

type Example struct {
	input  int
	output int
}

var personalData = 361527

var examplesPart1 = []Example{
	{input: 1, output: 0},
	{input: 12, output: 3},
	{input: 23, output: 2},
	{input: 1024, output: 31},
	{input: personalData, output: 326},
}

var examplesPart2 = []Example{
	{input: 1, output: 2},
	{input: 12, output: 23},
	{input: 23, output: 25},
	{input: personalData, output: 363010},
}

func TestSpiralMemoryPart1(t *testing.T) {
	for i := 0; i < len(examplesPart1); i++ {
		example := examplesPart1[i]
		res := SpiralMemoryPart1(example.input)
		if res != example.output {
			t.Error(fmt.Sprintf("Expected %d, but got %d", example.output, res))
		}
	}
}

func TestSpiralMemoryPart2(t *testing.T) {
	for i := 0; i < len(examplesPart2); i++ {
		example := examplesPart2[i]
		res := SpiralMemoryPart2(example.input)
		if res != example.output {
			t.Error(fmt.Sprintf("Expected %d, but got %d", example.output, res))
		}
	}
}
