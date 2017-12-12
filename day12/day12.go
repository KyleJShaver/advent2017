package day12

import (
	"strings"
	"strconv"
)

func Groups(input string) (int, int) {
	lines := strings.Split(input, "\n")
	network := make([]map[int]int, len(lines))
	for i := range network {
		network[i] = make(map[int]int)
	}
	for _, line := range lines {
		line = strings.Replace(line, " <-> ", ", ", 1)
		components := strings.Split(line, ", ")
		currentNumber, _ := strconv.Atoi(components[0])
		for i := 1; i < len(components); i++ {
			connection, _ := strconv.Atoi(components[i])
			network[currentNumber][connection]++
			network[connection][currentNumber]++
		}
	}
	touchable := make(map[int]int)
	traverseGraph(network, touchable, 0)
	touchableFromZero := len(touchable)
	seen := make([]int, len(lines))
	numGroups := 0
	for i := range seen {
		if seen[i] == 0 {
			groupTouchable := make(map[int]int)
			traverseGraph(network, groupTouchable, i)
			for j := range groupTouchable {
				seen[j]++
			}
			numGroups++
		}
	}
	return touchableFromZero, numGroups
}

func traverseGraph(network []map[int]int, touchable map[int]int, currPos int) {
	for i := range network[currPos] {
		if touchable[i] == 0 {
			touchable[i]++
			traverseGraph(network, touchable, i)
		}
	}
}