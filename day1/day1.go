package day1

func CaptchaPart1(input string) int32 {
	data := []byte(input)
	data = append(data, data[0]) // add first byte to the end to imitate a cycle
	var total int32 = 0
	for i := 0; i < len(data)-1; i++ {
		num := int32(data[i]) - 48 // ascii to int
		next := int32(data[i+1]) - 48
		if num == next {
			total += num
		}
	}
	return total
}

func CaptchaPart2(input string) int32 {
	data := []byte(input)
	step := len(data) / 2
	var total int32 = 0
	for i := 0; i < step; i++ {
		num := int32(data[i]) - 48 // ascii to int
		next := int32(data[i+step]) - 48
		if num == next {
			total += num + next
		}
	}
	return total
}