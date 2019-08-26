package main

import (
	"strings"
)

func main() {
	println(IsValid("{[]}"))
}

//IsValid - function for checking braces
func IsValid(s string) bool {
	if s == "" {
		return true
	}
	arr := strings.Split(s, "")
	if len(arr)%2 != 0 {
		return false
	}
	lastOpenBraces = []string{}
	for _, b := range arr {
		if !isAllowedSymbol(b) || !isCorrectBrace(b) {
			return false
		}
	}
	return len(lastOpenBraces) == 0
}

var lastOpenBraces []string
var braces = map[string]string{
	"}": "{",
	")": "(",
	"]": "[",
}

func isAllowedSymbol(b string) bool {
	return b == "{" || b == "}" || b == "(" || b == ")" || b == "[" || b == "]"
}

func isOpenBrace(b string) bool {
	return b == "[" || b == "{" || b == "("
}

func isCorrectBrace(b string) bool {
	if isOpenBrace(b) {
		lastOpenBraces = append(lastOpenBraces, b)
		return true
	}
	l := len(lastOpenBraces)
	if l == 0 {
		return false
	} else if lastOpenBraces[l-1] == braces[b] {
		lastOpenBraces = lastOpenBraces[:l-1]
		return true
	}
	return false
}
