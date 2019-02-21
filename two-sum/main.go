package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(TwoSum(nums, 17))
}

func TwoSum(nums []int, target int) []int {
	lent := len(nums)
	if lent == 0 {
		return []int{0, 0}
	}
	for i := 0; i < lent; i++ {
		for j := i + 1; j < lent; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
