package AddTwoNumbers

import (
	"testing"
)

type testCase struct {
	input1 *ListNode
	input2 *ListNode
	output *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	for _, c := range getTestCases() {
		res := AddTwoNumbers(c.input1, c.input2)
		if compareNodes(res, c.output) == false {
			t.Error("Failed! Nodes:", c.input1, c.input2, "res:", getStringFromNode(res), "should be:", getStringFromNode(c.output))
		}
	}
}

func compareNodes(n1 *ListNode, n2 *ListNode) bool {
	clone1 := copyNode(n1)
	clone2 := copyNode(n2)
	for {
		if clone1 == nil && clone2 == nil {
			return true
		} else if clone1 == nil || clone2 == nil {
			return false
		} else if clone1.Val != clone2.Val {
			return false
		}
		clone1 = copyNode(clone1.Next)
		clone2 = copyNode(clone2.Next)
	}
}

func getTestCases() []testCase {
	return []testCase{
		{
			createNode([]int{2, 4, 3}),
			createNode([]int{5, 6, 4}),
			createNode([]int{7, 0, 8}),
		},
		{
			createNode([]int{3, 4, 3}),
			createNode([]int{5, 6, 4}),
			createNode([]int{8, 0, 8}),
		},
		{
			createNode([]int{1, 2, 7}),
			createNode([]int{9, 6, 2}),
			createNode([]int{0, 9, 9}),
		},
		{
			createNode([]int{1, 1, 1}),
			createNode([]int{9, 9, 9}),
			createNode([]int{0, 1, 1, 1}),
		},
		{
			createNode([]int{1}),
			createNode([]int{9, 9, 9}),
			createNode([]int{0, 0, 0, 1}),
		},
		{
			createNode([]int{1}),
			createNode([]int{9, 9}),
			createNode([]int{0, 0, 1}),
		},
		{
			createNode([]int{1, 8}),
			createNode([]int{0}),
			createNode([]int{1, 8}),
		},
		{
			&ListNode{},
			createNode([]int{1, 8}),
			createNode([]int{1, 8}),
		},
		{
			createNode([]int{1, 8}),
			&ListNode{},
			createNode([]int{1, 8}),
		},
		{
			createNode([]int{0, 8, 6, 5, 6, 8, 3, 5, 7}),
			createNode([]int{6, 7, 8, 0, 8, 5, 8, 9, 7}),
			createNode([]int{6, 5, 5, 6, 4, 4, 2, 5, 5, 1}),
		},
		{
			createNode([]int{0, 8, 8, 8, 8, 2, 9, 3, 1, 1}),
			createNode([]int{0, 9, 1, 5, 5, 5, 1, 1, 6}),
			createNode([]int{0, 7, 0, 4, 4, 8, 0, 5, 7, 1}),
		},
		{
			&ListNode{},
			&ListNode{},
			&ListNode{},
		},
	}
}

func copyNode(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	b := *n
	return &b
}
