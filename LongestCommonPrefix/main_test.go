package LongestCommonPrefix

import "testing"

type testCase struct {
	input  []string
	output string
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, tCase := range getTestCasesLongestCommonPrefix() {
		result := LongestCommonPrefix(tCase.input)
		if result != tCase.output {
			t.Error("For strings", tCase.input, "prefix should be", tCase.input, "got", result)
		}
	}
}

func getTestCasesLongestCommonPrefix() []testCase {
	return []testCase{
		{[]string{"formula", "forl", "foma"}, "fo"},
		{[]string{}, ""},
		{[]string{"kaka"}, "kaka"},
		{[]string{"loal", "sdaf"}, ""},
		{[]string{","}, ","},
		{[]string{"", ""}, ""},
		{[]string{"c", "c"}, "c"},
		{[]string{"", "c"}, ""},
	}
}
