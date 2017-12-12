package day12

import (
	"fmt"
	"testing"
	"io/ioutil"
)

type Example struct {
	input string
	part1 int
	part2 int
}

var puzzleInputBytes, _ = ioutil.ReadFile("puzzle input.txt")
var puzzleInputStr = string(puzzleInputBytes)

var examples = []Example{
	{input: "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5", part1: 6, part2: 2},
	{input: puzzleInputStr, part1: 115, part2: 221},
}

func TestKnots(t *testing.T) {
	for i := 0; i < len(examples); i++ {
		example := examples[i]
		resPart1, resPart2 := Groups(example.input)
		if resPart1 != example.part1 || resPart2 != example.part2 {
			t.Error(fmt.Sprintf("Expected %d, %d but got %d, %d", example.part1, example.part2, resPart1, resPart2))
		}
	}
}
