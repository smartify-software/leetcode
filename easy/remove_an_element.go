package main

import "fmt"

func removeElement(nums []int, val int) int {
	// nums = [3,2,2,3], val = 3
	// start from index 0 and search for the element; if you find it then from that index on find another element
	// that's not the target element and swap
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := i; j < len(nums); j++ {
				if nums[j] != val {
					swap(nums, i, j)
					count++
				}
			}
		}
	}
	return count
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func main() {
	//count := removeElement([]int{3, 2, 2, 3}, 3)
	//nums := []int{3, 2, 2, 3}
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums := []int{2}
	count := removeElement(nums, 3)
	fmt.Printf("%+v, %d\n", nums, count)
}
