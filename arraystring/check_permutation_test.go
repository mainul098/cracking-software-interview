package arraystring

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	testCases := []struct {
		firstString  string
		secondString string
		result       bool
	}{
		{"hello", "elloh", true},
		{"mainul", "aminul", true},
		{"keya", "key", false},
		{"keya", "keye", false},
	}

	t.Logf("Should decide if two strings are permutation of the one another.")
	{
		for _, testCase := range testCases {
			t.Run(testCase.firstString, func(t *testing.T) {
				if r := checkPermutation(testCase.firstString, testCase.secondString); r != testCase.result {
					t.Errorf("Failed, Does not match with the expected result. Expected : %v, but Got : %v", testCase.result, r)
				} else {
					t.Logf("Success.")
				}
			})
		}
	}
}
