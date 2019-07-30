package main

import (
	"strconv"
	"strings"
)

func main() {
	println("hi")
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	arr := strings.Split(strconv.Itoa(x), "")
	length := len(arr)
	half := 0
	if length%2 == 0 {
		half = length / 2
	} else {
		half = (length - 1) / 2
	}
	for i := 0; i < half; i++ {
		if arr[i] != arr[length-i-1] {
			return false
		}
	}
	return true
}
