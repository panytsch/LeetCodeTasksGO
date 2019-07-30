package main

import "testing"

type testCase struct {
	input  int
	output bool
}

func TestIsPalindrome(t *testing.T) {
	for _, tCase := range getTestCasesIsPalindrome() {
		result := IsPalindrome(tCase.input)
		if result != tCase.output {
			t.Error("Wrong result for number", tCase.input, "Expected", tCase.output, "got", result)
		}
	}
}

func getTestCasesIsPalindrome() []testCase {
	return []testCase{
		{121, true},
		{11, true},
		{1, true},
		{122, false},
		{-122, false},
		{10, false},
		{101, true},
		{-101, false},
		{0, true},
	}
}
