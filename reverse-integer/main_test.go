package main

import "testing"

type task struct {
	input  int
	output int
}

var dataSet = []task{
	{23, 32},
	{0, 0},
	{5, 5},
	{-23, -32},
	{214474833647, 0},
	{-241474833647, 0},
	{-1534236469, 0},
	{1534236469, 0},
	{123, 321},
	{-123, -321},
}

func TestReverseInt(t *testing.T) {
	for _, data := range dataSet {
		res := ReverseInt(data.input)
		if res != data.output {
			t.Error("result", res, "isn't equal", data.output)
		}
	}
}
