package day10

import (
	"fmt"
	"testing"
)

type Example struct {
	input []int
	size  int
	part1 int
	part2 string
}

var personalData = []int{120, 93, 0, 90, 5, 80, 129, 74, 1, 165, 204, 255, 254, 2, 50, 113}

var examples = []Example{
	{input: []int{3, 4, 1, 5}, size: 5, part1: 12, part2: "04"},
	{input: personalData, size: 256, part1: 826, part2: "d067d3f14d07e09c2e7308c3926605c4"},
}

func TestKnots(t *testing.T) {
	for i := 0; i < len(examples); i++ {
		example := examples[i]
		resPart1, resPart2 := Knots(example.input, example.size)
		if resPart1 != example.part1 || resPart2 != example.part2 {
			t.Error(fmt.Sprintf("Expected %d, \"%s\" but got %d, \"%s\"", example.part1, example.part2, resPart1, resPart2))
		}
	}
}
