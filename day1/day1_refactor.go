package day1

func CaptchaRefactor(input string) (int32, int32) {
	var part1, part2 int32 = 0, 0
	length, step := len(input), len(input)/2
	for pos, ascii := range input {
		decVal := ascii - 48                           // ASCII to decimal value conversion
		nextVal := int32(input[(pos+1)%length] - 48)   // Next value in the captcha
		oppVal := int32(input[(pos+step)%length] - 48) // Value on the opposite position in the captcha
		if decVal == nextVal {
			part1 += decVal
		}
		if decVal == oppVal {
			part2 += decVal
		}
	}
	return part1, part2
}
