package day9

import (
	"regexp"
	"strings"
)

func Groups(input string) (int, int) {
	rx := regexp.MustCompile("!.")
	input = rx.ReplaceAllString(input, "")
	rx = regexp.MustCompile(",?<[\\{}'a-z\"<,]*>,?")
	garbageCount := 0
	foundGarbage := rx.FindAllString(input, -1)
	for _, match := range foundGarbage {
		match = strings.Trim(match, ",")
		garbageCount += len(strings.Trim(match, ",")) - 2
	}
	input = rx.ReplaceAllString(input, "")
	input = strings.Replace(input, ",", "", -1)
	groupCount, curr := 0, 0
	for _, char := range input {
		str := string(char)
		if str == "{" {
			curr += 1
			groupCount += curr
		} else if str == "}" {
			curr -= 1
		}
	}
	return groupCount, garbageCount
}
