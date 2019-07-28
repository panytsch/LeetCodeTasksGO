package main

import "testing"

type task struct {
	values []int
	result [][]int
}

var taskSet = []task{
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
	},
	{
		[]int{0, 0, 1},
		[][]int{
			{0, 0, 0},
		},
	},
	{
		[]int{0, 0},
		[][]int{
			{0, 0, 0},
		},
	},
	{
		[]int{},
		[][]int{
			{0, 0, 0},
		},
	},
}

func TestThreeSum(t *testing.T) {
	for _, task := range taskSet {
		v := ThreeSum(task.values)
		if len(v) != len(task.result) {
			t.Error(
				"For", task.values,
				"expected", task.result,
				"got", v,
			)
			continue
		}
		for k1, i := range v {
			for k2, j := range i {
				if j != task.result[k1][k2] {
					t.Error(
						"For", task.values,
						"expected", task.result,
						"got", v,
					)
				}
			}
		}
	}
}
