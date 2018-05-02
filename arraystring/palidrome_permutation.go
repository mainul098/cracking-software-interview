package arraystring

// Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palin­ drome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters.
// The palindrome does not need to be limited to just dictionary words.
// EXAMPLE
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco eta", etc.)

import (
	"strings"
)

func checkPalindromePermutation(text string) bool {
	// For empty string
	if len(text) == 0 {
		return true
	}

	// Convert to the lower to make case insensitive
	text = strings.ToLower(text)

	charCount := make(map[rune]int, 0)
	for _, ch := range text {
		// If space, skip
		if ch == ' ' {
			continue
		}

		// If we dont have the charecter in the set, add in the set
		if count := charCount[ch]; count == 0 {
			charCount[ch]++
		} else {
			// if already in the set,remove the char
			delete(charCount, ch)
		}
	}

	// Only one charecter can be odd numbered
	if len(charCount) > 2 {
		return false
	}

	return true
}
