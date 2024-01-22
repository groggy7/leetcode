package main

import (
	"fmt"
	"strings"
)

/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order,
starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
Example 2:

Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s
Example 3:

Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d
*/

func mergeAlternately(word1 string, word2 string) string {
	var merged strings.Builder
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		merged.WriteByte(word1[i])
		merged.WriteByte(word2[j])
		i++
		j++
	}

	if i < len(word1) {
		merged.WriteString(word1[i:])
	} else if j < len(word2) {
		merged.WriteString(word2[j:])
	}

	return merged.String()
}

func main() {
	fmt.Println(mergeAlternately("abc", "qrs"))
}
