package main

import (
	"fmt"
	"math"
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
	word1len := len(word1)
	word2len := len(word2)

	diff := math.Abs(float64(word1len) - float64(word2len))
	var index int
	var longWord string
	if word1len > word2len {
		index = word1len
		longWord = word1
	} else if word2len > word1len {
		index = word2len
		longWord = word2
	} else {
		index = word1len
	}

	var appendedString string
	for i := 0; i < index; i++ {
		appendedString += string(word1[i])
		appendedString += string(word2[i])
	}

	if diff == 0 {
		return appendedString
	} else {
		appendedString += longWord[len(longWord)-int(diff)-1:]
		return appendedString
	}
}

func main() {
	fmt.Println(mergeAlternately("abc", "qrs"))
}
