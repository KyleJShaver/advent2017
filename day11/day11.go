package day11

import (
	"strings"
)

func Steps(input string) (int, int) {
	steps := strings.Split(input, ",")
	finalDistance := calcDistance(steps)
	maxDistance := finalDistance
	for pos := range steps {
		currDist := calcDistance(steps[:len(steps) - 1 - pos])
		if currDist > maxDistance {
			maxDistance = currDist
		}
	}
	return finalDistance, maxDistance
}

func calcDistance(steps []string) int {
	stepCount := make(map[string]int)
	for _, step := range steps {
		stepCount[step]++
	}
	opposites := [][]string{{"s","n"}, {"ne","sw"}, {"se", "nw"}}
	reducers := [][]string{{"ne", "nw", "n"}, {"se", "sw", "s"}, {"sw", "n", "nw"}, {"se", "n", "nw"}, {"nw", "s", "sw"}, {"ne", "s", "se"}}
	for _, opposite := range opposites {
		min := stepCount[opposite[0]]
		if stepCount[opposite[1]] < min {
			min = stepCount[opposite[1]]
		}
		stepCount[opposite[0]] -= min
		stepCount[opposite[1]] -= min
	}
	for _, reducer := range reducers {
		min := stepCount[reducer[0]]
		if stepCount[reducer[1]] < min {
			min = stepCount[reducer[1]]
		}
		stepCount[reducer[0]] -= min
		stepCount[reducer[1]] -= min
		stepCount[reducer[2]] += min
	}
	total := 0
	for _, opp := range opposites {
		total += stepCount[opp[0]]
		total += stepCount[opp[1]]
	}
	return total
}