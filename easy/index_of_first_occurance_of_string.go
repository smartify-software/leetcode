package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if isIn(haystack[i:], needle) {
			return i
		}
	}
	return -1
}

func isIn(s string, needle string) bool {
	if len(needle) > len(s) {
		return false
	}
	for i := 0; i < len(needle); i++ {
		// invariant here is that
		// s should be at least as big as needle
		// for each value of needle it should equal to that value of s
		if s[i] != needle[i] {
			return false
		}
	}
	return true
}

func main() {
	i := strStr("mississippi", "issippi")
	fmt.Printf("%d\n", i)
}
