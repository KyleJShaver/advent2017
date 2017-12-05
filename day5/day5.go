package day5

func Trampolines(input []int) (int, int) {
	parts := []int{0, 0}
	inputs := [][]int{input, make([]int, len(input))}
	for pos, val := range input {
		inputs[1][pos] = val
	}
	for pos, nums := range inputs {
		instruction := 0
		for instruction >= 0 && instruction < len(input) {
			curr := nums[instruction]
			increase := 1
			if curr >= 3 && pos == 1 {
				increase = -1
			}
			nums[instruction] += increase
			instruction += curr
			parts[pos]++
		}
	}
	return parts[0], parts[1]
}

