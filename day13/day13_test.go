package day13

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type Example struct {
	input string
	part1 int
	part2 int
}

var puzzleInputBytes, _ = ioutil.ReadFile("puzzle input.txt")
var puzzleInputStr = string(puzzleInputBytes)

var examples = []Example{
	{input: "0: 3\n1: 2\n4: 4\n6: 4", part1: 24, part2: 10},
	{input: puzzleInputStr, part1: 2264, part2: 3875838},
}

func TestFirewall(t *testing.T) {
	for i := 0; i < len(examples); i++ {
		example := examples[i]
		resPart1, resPart2 := Firewall(example.input)
		if resPart1 != example.part1 || resPart2 != example.part2 {
			t.Error(fmt.Sprintf("Expected %d, %d but got %d, %d", example.part1, example.part2, resPart1, resPart2))
		}
	}
}
