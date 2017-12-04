package day2

import "math"

func ChecksumRefactor(input [][]int) (int, int) {
	part1, part2, smallest, largest, multiple, factor := 0, 0, 0, 0, 0, 0
	for _, row := range input {
		smallest = int(math.MaxInt32)
		largest = int(math.MinInt32)
		for colPos, currVal := range row {
			if currVal < smallest {
				smallest = currVal
			}
			if currVal > largest {
				largest = currVal
			}
			for i := colPos + 1; i < len(row); i++ {
				checkVal := row[i]
				if checkVal < currVal {
					factor = checkVal
					multiple = currVal
				} else {
					factor = currVal
					multiple = checkVal
				}
				if multiple%factor == 0 {
					part2 += multiple / factor
				}
			}
		}
		part1 += largest - smallest
	}
	return part1, part2
}
