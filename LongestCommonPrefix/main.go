package LongestCommonPrefix

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	} else if length == 1 {
		return strs[0]
	}
	result := ""
	i := 0
	l := len(strs[0])
	for {
		for j := 1; j < length; j++ {
			if len(strs[j]) <= i || l <= i {
				return result
			}
			if strs[j][i] != strs[j-1][i] || strs[j][i] != strs[0][i] {
				return result
			}
		}
		result += string(strs[0][i])
		i++
	}
}
