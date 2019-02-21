package main

import "testing"

type task struct {
	slice  []int
	target int
	result []int
}

var entry = []task{
	{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 7, 11, 15}, 18, []int{1, 2}},
	{[]int{2, 7, 11, 15}, 22, []int{1, 3}},
	{[]int{2, 7, 11, 15}, 26, []int{2, 3}},
	{[]int{2, 7, 11, 15}, 30, []int{0, 0}},
	{[]int{}, 30, []int{0, 0}},
}

func TestTwoSum(t *testing.T) {
	for _, task := range entry {
		v := TwoSum(task.slice, task.target)
		if v[1] != task.result[1] && v[0] != task.result[0] {
			t.Error(
				"For", task.slice,
				"expected", task.result,
				"got", v,
			)
		}
	}
}
