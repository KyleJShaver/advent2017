package day4

import (
	"strings"
	"sort"
)

func ValidPassphrases(input string) (int, int) {
	lines := strings.Split(input, "\n")
	part1, part2 := 0, 0
	for _, line := range lines {
		distinct, anagram := make(map[string]int), make(map[string]int)
		distinctValid, anagramValid := true, true
		words := strings.Split(line, " ")
		for _, word := range words {
			if distinct[word] == 0 {
				distinct[word] = 1
			} else {
				distinctValid = false
			}
			wordChars := strings.Split(word, "")
			sort.Strings(wordChars)
			alphabetized := strings.Join(wordChars, "")
			if anagram[alphabetized] == 0 {
				anagram[alphabetized] = 1
			} else {
				anagramValid = false
			}
		}
		if distinctValid {
			part1++
		}
		if anagramValid {
			part2++
		}
	}
	return part1, part2
}