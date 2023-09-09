package main

import "fmt"

func searchInsert(nums []int, target int) int {
	//  nums = [1,3,5,6], target = 2
	// do binary search ; if target >  nums[i/2] then search left else search right
	// if you find search then return i
	// if you can't f
	return binarySearch(nums, target)
}

func BSearch(nums []int, target int) int {
	mid := len(nums) / 2
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	if target == nums[mid] {
		return mid
	}
	if target < nums[mid] {
		return BSearch(nums[:mid], target)
	} else {
		return BSearch(nums[mid+1:], target)
	}
}

func binarySearch(nums []int, i int) int {
	mid := len(nums) / 2
	if i > len(nums) {
		return -1
	}
	if i == nums[mid] {
		return mid
	}
	if i < nums[mid] {
		binarySearch(nums[mid+1:], mid)
	} else {
		binarySearch(nums[:mid], mid)
	}
	return -1
}

func main() {
	index := searchInsert([]int{1, 3, 5, 6}, 3)
	fmt.Printf("%d\n", index)
}
