package main

import (
	"testing"
)

type testCase struct {
	in  string
	out bool
}

func getCases() []testCase {
	return []testCase{
		{"{}", true},
		{"{[]}", true},
		{"[{}]", true},
		{"", true},
		{"{([])}", true},
		{"()", true},
		{"[][]{}()", true},
		{"([{}])", true},
		{"{)", false},
		{"{{", false},
		{"{[)", false},
		{"({)}", false},
		{"[{])", false},
		{"{([})}", false},
		{"{b}", false},
	}
}

func TestIsValid(t *testing.T) {
	for _, task := range getCases() {
		v := IsValid(task.in)
		// log.Println(task.in, task.out, "result:", v)
		if v != task.out {
			t.Error(
				"For", task.in,
				"expected", task.out,
				"got", v,
			)
		}
	}
}
