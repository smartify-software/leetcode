package main

import "fmt"

func removeDuplicates(nums []int) int {
	// 1, 1, 1, 2, 2, 3
	// 1, 2, 1, 2, 2, 3
	// 1, 2, 3

	// iter nums using i, j
	// j will find next new number and i++ will store it
	// while j < len(nums)
	count := 0
	i := 0
	j := 0
	for {
		for j+1 < len(nums) && nums[j] == nums[j+1] {
			j++
		}
		j++
		nums[i+1] = nums[j]
		i++
		count++
		if j > len(nums)-1 {
			return count
		}
	}
}

func main() {
	ints := []int{1, 1, 1, 1, 1, 2, 2, 3, 4}
	duplicates := removeDuplicates(ints)
	fmt.Printf("%+v", ints)
	fmt.Printf("%d", duplicates)
}
