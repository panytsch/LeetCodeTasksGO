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
	s1 := convSliceToInt(reverseString(strings.Split(getStringFromNode(l1), "")))
	s2 := convSliceToInt(reverseString(strings.Split(getStringFromNode(l2), "")))
	log.Println(s1, s2)
	return createNode(addTwoSlices(s1, s2))
}

func addTwoSlices(s1 []int, s2 []int) []int {
	more9 := false
	var res []int
	cond := len(s1) > len(s2)
	var big []int
	var small []int
	if cond {
		big = s1
		small = s2
	} else {
		big = s2
		small = s1
	}
	lenSmall := len(small) - 1
	lenBig := len(big) - 1
	j := lenSmall
	for i := lenBig; i >= 0; i-- {
		if j < 0 {
			summ := big[i]
			if more9 {
				summ++
				if summ > 9 {
					summ = summ % 10
				} else {
					more9 = false
				}
			}
			res = append(res, summ)
		} else {
			summ := big[i] + small[j]
			if more9 {
				summ++
				more9 = false
			}
			if summ > 9 {
				more9 = true
				summ = summ % 10
			}
			res = append(res, summ)
		}
		j--
	}
	if more9 {
		res = append(res, 1)
	}
	return res
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
