package day2

import "math"

func ChecksumPart1(input [][]int) int {
	total, smallest, largest := 0, 0, 0
	for i := 0; i < len(input); i++ {
		row := input[i]
		smallest = int(math.MaxInt32)
		largest = int(math.MinInt32)
		for j := 0; j < len(row); j++ {
			num := row[j]
			if num < smallest {
				smallest = num
			}
			if num > largest {
				largest = num
			}
		}
		total += largest - smallest
	}
	return total
}

func ChecksumPart2(input [][]int) int {
	total, multiple, factor := 0, 0, 0
	for i := 0; i < len(input); i++ {
		row := input[i]
		for j := 0; j < len(row)-1; j++ {
			for k := j + 1; k < len(row); k++ {
				if row[k] < row[j] {
					factor = row[k]
					multiple = row[j]
				} else {
					factor = row[j]
					multiple = row[k]
				}
				if multiple%factor == 0 {
					k = len(row)
					j = k
					total += multiple / factor
				}
			}
		}
	}
	return total
}
