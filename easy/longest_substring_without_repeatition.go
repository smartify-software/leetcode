package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	for i, _ := range s {
		return math.Max(1 + , lengthOfLongestSubstring(s[i+1:]))
	}
	return 0
}
