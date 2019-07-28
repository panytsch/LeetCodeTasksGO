package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {

}

func ReverseInt(x int) int {
	if x > -1 && x < 10 {
		return x
	}
	sX := strings.Split(strconv.Itoa(x), "")
	cond := sX[0] == "-"
	length := len(sX)
	for i := length/2 - 1; i >= 0; i-- {
		opp := length - 1 - i
		sX[i], sX[opp] = sX[opp], sX[i]
	}
	if cond {
		sX = append([]string{"-"}, sX...)
		sX[length] = ""
	}
	s, e := strconv.Atoi(strings.Join(sX, ""))
	if e != nil || s > math.MaxInt32 || s < math.MinInt32 {
		return 0
	}
	return s
}
