package arraystring

import "testing"

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{"hello", false},
		{"hellolleh", true},
		{"Tact Coa", true},
		{"", true},
	}

	t.Logf("Should check the string and check the permutattion of the charecters is a plindrome")
	{
		for _, testCase := range testCases {
			t.Run(testCase.input, func(t *testing.T) {
				if r := checkPalindromePermutation(testCase.input); r != testCase.output {
					t.Errorf("Failed, Does not match with the expected result. Expected : %v, but Got : %v", testCase.output, r)
				} else {
					t.Logf("Success.")
				}
			})
		}
	}
}
