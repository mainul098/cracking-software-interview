package arraystring

// Implement an algorithm to determine if a string has all unique characters.
// TimeComplexity : O(n)
func checkUnique(text string) bool {
	charSet := make(map[rune]bool, 0)
	for _, ch := range text {
		if _, ok := charSet[ch]; ok == true {
			return false
		}

		charSet[ch] = true
	}

	return true
}

// Follwoup question : What if you cannot use additional data structures?
// TimeComplexity : O(n^2)
func checkUniqueUsingNoAdditionalSpace(text string) bool {
	for i := 0; i < len(text); i++ {
		for j := i + 1; j < len(text); j++ {
			if text[i] == text[j] {
				return false
			}
		}
	}
	return true
}
