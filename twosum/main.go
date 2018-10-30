package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 8, 6, 7}
	fmt.Println(twoSum(nums, 10))
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
