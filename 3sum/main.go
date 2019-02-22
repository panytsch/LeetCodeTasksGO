package main

import "fmt"

func main() {
	data := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(data)
	fmt.Println("result: ", ThreeSum(data))
}

func ThreeSum(nums []int) [][]int {
	lent := len(nums)
	if lent < 3 {
		return [][]int{{0, 0, 0}}
	}
	//main logic start
	var result [][]int

	for i := 0; i < lent/2; i++ {
		for j := lent - 1; j > lent/2; j-- {
			for k := i + 1; k < j; k++ {
				fmt.Println(i, j, k)
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	if len(result) == 0 {
		return [][]int{{0, 0, 0}}
	}
	return result
}
