package main

import "fmt"

func main() {
	data := []int{-5, 1, -10, 2, -7, -13, -3, -8, 2, -15, 9, -3, -15, 13, -6, -10, 5, 6, 11, 1, 13, -12, 14, 6, 11, 4, 13, -14, 0, 11, 1, 10, -11, 6, -11, -10, 8, 2, -3, -13, -6, -9, -9, -6, 11, -8, -9, 1, 13, 9, 9, 3, 13, 0, -6, 1, -10, -15, 3, 5, 3, 11, -8, 0, 2, -11, 5, -13, 6, 9, -11, 7, 8, -13, 8, 4, -6, 14, 13, -15, 1, 7, -5, -1, -7, 5, 7, -2, -3, -13, 10, 7, 13, 9, -8, -8, 13, 12, -6, 4, 7, -10, -12, -8, -8, 11, 11, -6, 3, 9, -14, -11, 2, -4, -5, 10, 8, -13, -7, 12, -10, 10}
	fmt.Println(data)
	fmt.Println("result: ", ThreeSum(data))
}

func ThreeSum(nums []int) [][]int {
	lent := len(nums)
	if lent < 3 {
		return [][]int{{0, 0, 0}}
	}
	var result [][]int
	//own sort
	for k1 := 0; k1 < lent-1; k1++ {
		min1 := nums[k1]
		for k2 := k1 + 1; k2 < lent; k2++ {
			if min1 > nums[k2] {
				nums[k1], nums[k2] = nums[k2], nums[k1]
			}
		}
	}
	compare := func(n, m []int) bool {
		return (n[0] == m[0] || n[0] == m[1] || n[0] == m[2]) &&
			(n[1] == m[0] || n[1] == m[1] || n[1] == m[2]) &&
			(n[2] == m[0] || n[2] == m[1] || n[2] == m[2])
	}

	lenRes := 0
	for i := 0; i < lent-2; i++ {
		for j := i + 1; j < lent-1; j++ {
			for k := j + 1; k < lent; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					newRes := []int{nums[i], nums[j], nums[k]}
					if lenRes != 0 && compare(newRes, result[lenRes-1]) {
						continue
					}
					result = append(result, newRes)
					lenRes++
				}
			}
		}
	}
	if len(result) == 0 {
		return [][]int{{0, 0, 0}}
	}
	return result
}
