package arraystring

// Given two strings,write a method to decide if one is a permutation of the other.
// Solution: check the count of charecter for each strings and see if the count match
func checkPermutation(firstText, secondText string) bool {
	if len(firstText) != len(secondText) {
		return false
	}

	charMap := make(map[rune]int, 0)
	for _, ch := range firstText {
		charMap[ch]++
	}

	for _, ch := range secondText {
		charMap[ch]--
		if charMap[ch] < 0 {
			return false
		}
	}
	return true
}
