package arraystring

import (
	"testing"
)

func TestCheckUnique(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{"hello", false},
		{"helo", true},
		{"", true},
	}

	t.Logf("Should check the string and validate the uniqueness of the charecters")
	{
		for _, testCase := range testCases {
			t.Run(testCase.input, func(t *testing.T) {
				if r := checkUnique(testCase.input); r != testCase.output {
					t.Errorf("Failed, Does not match with the expected result. Expected : %v, but Got : %v", testCase.output, r)
				} else {
					t.Logf("Success.")
				}
			})
		}
	}
}

func TestCheckUniqueUsingNoAdditionalSpace(t *testing.T) {

	testCases := []struct {
		input  string
		output bool
	}{
		{"hello", false},
		{"helo", true},
		{"", true},
	}

	t.Logf("Should check the string and validate the uniqueness of the charecters")
	{
		for _, testCase := range testCases {
			t.Run(testCase.input, func(t *testing.T) {
				if r := checkUniqueUsingNoAdditionalSpace(testCase.input); r != testCase.output {
					t.Errorf("Failed, Does not match with the expected result for '%v'. Expected : %v, but Got : %v", testCase.input, testCase.output, r)
				} else {
					t.Logf("Success.")
				}
			})
		}
	}
}
