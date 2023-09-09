package main

import "fmt"

func longestCommonPrefix(strings []string) string {
	// consider strings[0] as reference
	// for i as index in string[0] if string[0][i] is  equal to other strings then increment i
	// return string[0][:i]
	for i := 0; i < len(strings[0]); i++ {
		for j := 0; j < len(strings); j++ {
			if i >= len(strings[j]) {
				return strings[0][:i]
			}
			if strings[0][i] != strings[j][i] {
				return strings[0][:i]
			}
		}
	}
	return strings[0]
}

func main() {
	strings := []string{"flower", "flow", "flight"}
	//strings := []string{"race", "racecar", "rac"}
	out := longestCommonPrefix(strings)
	fmt.Printf("%+v\n", out)
}
