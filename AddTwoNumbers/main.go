package AddTwoNumbers

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l2.Val == 0 {
		return &ListNode{}
	} else if l1.Val == 0 {
		return createNode(convSliceToInt(convIntToSlice(convertToNumber(l2))))
	} else if l2.Val == 0 {
		return createNode(convSliceToInt(convIntToSlice(convertToNumber(l1))))
	}
	sum := rInt(convertToNumber(l1)) + rInt(convertToNumber(l2))
	return createNode(reverseInt(convSliceToInt(convIntToSlice(sum))))
}

//func addTwoNodes(l1 *ListNode, l2 *ListNode) []int {
//	var res []int
//	more9 := false
//	for {
//		sum := 0
//		if l1 == nil && l2 == nil {
//			return res
//		} else if l1 == nil {
//			sum = l2.Val
//		} else if l2 == nil {
//			sum = l1.Val
//		} else {
//			sum = l1.Val + l2.Val
//		}
//
//		if more9 == true {
//			sum ++
//			more9 = false
//		}
//
//		if sum > 9 {
//			sum = sum % 10
//			more9 = true
//		}
//
//		res = append(res, sum)
//
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//}

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

func rInt(x int) int {
	if x > -1 && x < 10 {
		return x
	}
	sX := strings.Split(strconv.Itoa(x), "")
	length := len(sX)
	for i := length/2 - 1; i >= 0; i-- {
		opp := length - 1 - i
		sX[i], sX[opp] = sX[opp], sX[i]
	}
	s, e := strconv.Atoi(strings.Join(sX, ""))
	if e != nil {
		return 0
	}
	return s
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
