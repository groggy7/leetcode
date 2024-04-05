/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different
word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
*/

function validAnagram(s, t) {
    if(s.length != t.length) {
        return false
    }

    let charCounts = {}

    for(i = 0; i < s.length; i++) {
        charCounts[s[i]] = (charCounts[s[i]] || 0) + 1
    }

    for(i = 0; i < t.length; i++) {
        if (!charCounts[t[i]]) {
            return false; 
        }
        charCounts[t[i]]--;
    }

    return true
}

validAnagram("abcdef", "a")