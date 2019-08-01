package PowerOfTwo

func isPowerOfTwo(n int) bool {
	if (n%2 == 0) && (n != 0) {
		return isPowerOfTwo(n / 2)
	} else if n == 1 {
		return true
	} else {
		return false
	}
}
