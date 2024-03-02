package main

import (
	"fmt"
	"sort"
)

// Solution 1
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runes := []rune(s)
	runet := []rune(t)

	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	sort.Slice(runet, func(i, j int) bool { return runet[i] < runet[j] })

	return string(runes) == string(runet)
}

// This is a good approach but sort operations will take likely O(nlogn) time complexity
// So its a heavy operation, instead we can do something like that:
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCounts := make(map[rune]int)
	for _, char := range s {
		charCounts[char]++
	}

	for _, char := range t {
		charCounts[char]--
		if charCounts[char] < 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isAnagram2("anag", "gana"))
}
