package day10

import (
	"fmt"
	"strings"
)

func Knots(input []int, size int) (int, string) {
	listA, listB := make([]int, size), make([]int, size)
	for i := range listA {
		listA[i] = i
		listB[i] = i
	}
	pos, skip := 0, 0
	for _, val := range input {
		pos, skip = inversion(listA, val, pos, skip)
	}
	part1 := listA[0] * listA[1]
	inputStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ","), "[]")
	inputB := make([]int, len(inputStr))
	for i, val := range inputStr {
		inputB[i] = int(val)
	}
	inputB = append(inputB, []int{17, 31, 73, 47, 23}...)
	pos, skip = 0, 0
	iterations := 64
	for i := 0; i < iterations; i++ {
		for _, val := range inputB {
			pos, skip = inversion(listB, val, pos, skip)
		}
	}
	outputStr := ""
	for i := 0; i < size; i += 16 {
		hashInt := listB[i]
		for j := 1; j < 16 && j < size; j++ {
			hashInt = hashInt ^ listB[i+j]
		}
		addStr := fmt.Sprintf("%x", hashInt)
		if len(addStr) == 1 {
			addStr = fmt.Sprintf("0%s", addStr)
		}
		outputStr = fmt.Sprintf("%s%s", outputStr, addStr)
	}
	return part1, outputStr
}

func inversion(list []int, length, pos, skip int) (int, int) {
	reversed := make([]int, length)
	size := len(list)
	for i := 0; i < length; i++ {
		reversed[length-1-i] = list[(pos+i)%size]
	}
	for i := 0; i < length; i++ {
		list[(pos+i)%size] = reversed[i]
	}
	pos = (pos + length + skip) % size
	skip += 1
	return pos, skip
}
