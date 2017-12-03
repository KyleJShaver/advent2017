package day3

import (
	"math"
)

func SpiralMemoryPart1(input int) int {
	pos := []int{0, 0}
	size := []int{1, 1}
	step := 0
	increase := true
	delta := -1
	degrees := 0
	traveled := 0
	for traveled <= 2*input {
		if increase {
			step += 1
		}
		increase = !increase
		stepDelta := 0
		if traveled+step >= input-1 {
			delta = traveled + 1 + step - input
			stepDelta = delta
		}
		switch degrees {
		case 0:
			pos[0] += step - stepDelta
			size[0] += 1
			break
		case 90:
			pos[1] = 0 + stepDelta
			size[1] += 1
			break
		case 180:
			pos[0] = 0 + stepDelta
			size[0] += 1
			break
		case 270:
			pos[1] += step - stepDelta
			size[1] += 1
			break
		}
		if delta >= 0 {
			centerX := math.Floor(float64((size[0] - 1) / 2))
			centerY := math.Floor(float64((size[1]) / 2))
			posX := float64(pos[0])
			posY := float64(pos[1])
			distance := math.Abs(centerX-posX) + math.Abs(centerY-posY)
			return int(distance)
		}
		degrees = (degrees + 90) % 360
		traveled += step
	}
	return 0
}

func SpiralMemoryPart2(input int) int {
	pos := []int{0, 0}
	size := []int{1, 1}
	step := 0
	increase := true
	delta := -1
	degrees := 0
	traveled := 0
	for traveled <= 2*input {
		if increase {
			step += 1
		}
		increase = !increase
		stepDelta := 0
		if traveled+step >= input-1 {
			delta = traveled + 1 + step - input
			stepDelta = delta
		}
		switch degrees {
		case 0:
			pos[0] += step - stepDelta
			size[0] += 1
			break
		case 90:
			pos[1] = 0 + stepDelta
			size[1] += 1
			break
		case 180:
			pos[0] = 0 + stepDelta
			size[0] += 1
			break
		case 270:
			pos[1] += step - stepDelta
			size[1] += 1
			break
		}
		if delta >= 0 {
			size[0] += 6
			size[1] += 6
			return memoryWithSize(size, input)
		}
		degrees = (degrees + 90) % 360
		traveled += step
	}
	return 0
}

func memoryWithSize(size []int, upTo int) int {
	memory := make([][]int, size[1])
	for i := range memory {
		memory[i] = make([]int, size[0])
	}
	centerX := int(math.Floor(float64((size[0] - 1) / 2)))
	centerY := int(math.Floor(float64((size[1]) / 2)))
	pos := []int{centerX, centerY}
	memory[pos[0]][pos[1]] = 1
	step := 0
	increase := true
	degrees := 0
	traveled := 1
	for i := 0; i < size[0]*size[1]; i++ {
		if increase {
			step += 1
		}
		increase = !increase
		for stepIncrementer := 0; stepIncrementer < step; stepIncrementer++ {
			switch degrees {
			case 0:
				pos[0] += 1
				break
			case 90:
				pos[1] -= 1
				break
			case 180:
				pos[0] -= 1
				break
			case 270:
				pos[1] += 1
				break
			}
			for j := 1; j >= -1 && pos[0]+j >= 0; j-- {
				for k := 1; k >= -1 && pos[1]+k >= 0; k-- {
					if j == 0 && k == 0 {
						continue
					} else {
						memory[pos[0]][pos[1]] += memory[pos[0]+j][pos[1]+k]
					}
				}
			}
			traveled++
			memVal := memory[pos[0]][pos[1]]
			if memVal > upTo {
				return memVal
			}
		}
		degrees = (degrees + 90) % 360
	}
	return 0
}
