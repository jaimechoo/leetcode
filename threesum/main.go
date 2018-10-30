package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	var container [][]int
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					container = append(container, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return container
}

func exsit(arrys [][]int, nums []int) {
	for _, arry := range arrys {
		for _, _ := range arry {

		}
	}
}
