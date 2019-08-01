package PowerOfTwo

import "testing"

type testCase struct {
	input  int
	output bool
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, c := range getTestCases() {
		res := isPowerOfTwo(c.input)
		if res != c.output {
			t.Error(c.input, "Expect:", c.output, "got:", res)
		}
	}
}

func getTestCases() []testCase {
	return []testCase{
		{2, true},
		{4, true},
		{0, false},
		{1, true},
		{16, true},
		{218, false},
	}
}
