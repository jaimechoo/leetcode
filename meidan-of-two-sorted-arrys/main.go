package main

import (
	"fmt"
)

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。

//请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。

//示例 1:
//nums1 = [1, 3]
//nums2 = [2]
//中位数是 2.0

//示例 2:
//nums1 = [1, 2]
//nums2 = [3, 4]
//中位数是 (2 + 3)/2 = 2.5

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	median := findMedianSortedArrays(nums1, nums2)

	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)
	fmt.Println("median:", median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	var lens = (len(nums1) + len(nums2))
	var a int = 0
	var b int = 0
	for a < len(nums1) || b < len(nums2) {
		if a < len(nums1) && b < len(nums2) {
			if nums1[a] <= nums2[b] {
				nums = append(nums, nums1[a])
				a++
			} else {
				nums = append(nums, nums2[b])
				b++
			}
		} else {
			if a < len(nums1) {
				nums = append(nums, nums1[a])
				a++
			} else {
				nums = append(nums, nums2[b])
				b++
			}
		}
	}
	if lens%2 == 0 {
		return float64(nums[lens/2-1]+nums[lens/2]) / 2
	} else {
		return float64(nums[lens/2])
	}
}
