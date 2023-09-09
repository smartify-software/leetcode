package main

import "fmt"

func lengthOfLastWord(s string) int {
	// if we encounter is a space the reset length
	// store the current and last word count
	// if we are at a space and this is the end of the string then return last
	// if we are at a space and this is the end of the string then return curr
	count := 0
	last := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && s[i] == ' ' {
			return last
		}
		if s[i] == ' ' {
			count = 0
		} else {
			count++
			last = count
		}
	}
	return count
}

func main() {
	//word := lengthOfLastWord("fly me   to   the moon    ")
	//word := lengthOfLastWord("Hello World")
	word := lengthOfLastWord("luffy is still joyboy")
	fmt.Printf("%d\n", word)
}
