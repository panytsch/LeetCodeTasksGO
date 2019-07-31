package AddTwoNumbers

import (
	"log"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l2.Val == 0 && l1.Next == nil && l2.Next == nil {
		return &ListNode{}
	} else if l1.Val == 0 && l1.Next == nil {
		return createNode(convSliceToInt(convIntToSlice(convertToNumber(l2))))
	} else if l2.Val == 0 && l2.Next == nil {
		return createNode(convSliceToInt(convIntToSlice(convertToNumber(l1))))
	}
	s1 := makeIntFromSlice(reverseString(strings.Split(getStringFromNode(l1), "")))
	s2 := makeIntFromSlice(reverseString(strings.Split(getStringFromNode(l2), "")))
	log.Println(s1, s2) //todo convert 1000000000000000000000000000001 in to 0
	return createNode(reverseInt(convSliceToInt(convIntToSlice(s1 + s2))))
}

func makeIntFromSlice(i []string) int {
	s, e := strconv.Atoi(strings.Join(i, ""))
	if e != nil {
		return 0
	}
	return s
}

func convertToNumber(node *ListNode) int {
	if node == nil {
		return 0
	}
	res, err := strconv.Atoi(getStringFromNode(node))
	if err != nil {
		return 0
	}
	return res
}

func getStringFromNode(node *ListNode) string {
	if node == nil {
		return ""
	} else if node.Next == nil {
		return strconv.Itoa(node.Val)
	} else {
		return strconv.Itoa(node.Val) + getStringFromNode(node.Next)
	}
}

func createNode(in []int) *ListNode {
	if len(in) == 0 {
		return nil
	}
	node := &ListNode{
		Val:  in[0],
		Next: nil,
	}
	in = remove(in, 0)
	node.Next = createNode(in)
	return node
}

func remove(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func reverseInt(x []int) []int {
	l := len(x)
	if l < 2 {
		return x
	}
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - 1 - i
		x[i], x[opp] = x[opp], x[i]
	}
	return x
}

func reverseString(x []string) []string {
	l := len(x)
	if l < 2 {
		return x
	}
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - 1 - i
		x[i], x[opp] = x[opp], x[i]
	}
	return x
}

func convSliceToInt(sl []string) []int {
	ss := make([]int, len(sl))
	for i, v := range sl {
		ss[i], _ = strconv.Atoi(v)
	}
	return ss
}

func convIntToSlice(i int) []string {
	return strings.Split(strconv.Itoa(i), "")
}
