package _rune

import "github.com/ekwok1/aoc-2021/utilities/slice"

func IsVowel(r rune) bool {
	vowels := []interface{}{'a', 'e', 'i', 'o', 'u'}
	return slice.Contains(&vowels, r)
}
