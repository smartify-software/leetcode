package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5}
	target := 9
	numbers := SumNumbers(nums, target)
	fmt.Printf("numbers: %v\n", numbers)
}

func SumNumbers(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
