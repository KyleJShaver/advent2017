package day13

import (
	"math"
	"strconv"
	"strings"
)

func Firewall(input string) (int, int) {
	lines := strings.Split(input, "\n")
	scannerSizes := make(map[int]int)
	maxDepth := 0
	for _, line := range lines {
		components := strings.Split(line, ": ")
		layerDepth, _ := strconv.Atoi(components[0])
		layerRange, _ := strconv.Atoi(components[1])
		scannerSizes[layerDepth] = layerRange
		maxDepth = layerDepth
	}
	firstSeverity, severity := -1, 0
	scannerPositions := make(map[int]int)
	for depth := range scannerSizes {
		scannerPositions[depth] = 0
	}
	for packetDepth := 0; packetDepth <= maxDepth; packetDepth++ {
		if scannerPositions[packetDepth] == 0 && scannerSizes[packetDepth] > 0 {
			severity += scannerSizes[packetDepth] * packetDepth
			if packetDepth == maxDepth {
				break
			}
		}
		for depth, currPos := range scannerPositions {
			if currPos == scannerSizes[depth]-1 {
				scannerPositions[depth] *= -1
			}
			scannerPositions[depth] += 1
		}
	}
	if firstSeverity < 0 {
		firstSeverity = severity
	}
	severity = 0
	safeDelay := 0
	for safeDelay = 1; safeDelay < math.MaxInt32; safeDelay++ {
		isSafe := true
		for depth, size := range scannerSizes {
			if (depth+safeDelay)%(2*(size-1)) == 0 {
				isSafe = false
				break
			}
		}
		if isSafe {
			return firstSeverity, safeDelay
		}
	}
	return firstSeverity, 0
}
