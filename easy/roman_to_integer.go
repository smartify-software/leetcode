package main

import "fmt"

func main() {
	toInt := romanToInt("MCMXCIV")
	fmt.Printf("toInt: %v\n", toInt)
}

func romanToInt(str string) int {
	store := make(map[string]int)
	store["M"] = 1000
	store["D"] = 500
	store["C"] = 100
	store["L"] = 50
	store["X"] = 10
	store["V"] = 5
	store["I"] = 1
	val := 0
	i := 0
	for {
		if i >= len(str)-1 {
			break
		}
		if store[string(str[i])] < store[string(str[i+1])] {
			val += store[string(str[i+1])] - store[string(str[i])]
			i += 2
		} else {
			val += store[string(str[i])]
			i += 1
		}
	}
	if i == len(str)-1 {
		val += store[string(str[i])]
	}
	return val
}
