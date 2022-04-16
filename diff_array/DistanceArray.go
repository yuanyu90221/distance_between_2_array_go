package diff_array

import (
	"fmt"
	"sort"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func findDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	fmt.Println(arr1, arr2)
	result := len(arr1)
	for _, val := range arr1 {
		small := FindCloseElement(arr2, val)
		fmt.Println(val, arr2[small])
		if abs(val-arr2[small]) <= d {
			result -= 1
		}
	}
	return result
}

func FindCloseElement(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for right > left {
		mid := (right + left) / 2
		if target < nums[mid] {
			right = mid - 1
		}
		if target > nums[mid] {
			left = mid + 1
		}
		if target == nums[mid] {
			return mid
		}
	}
	mid := (right + left) / 2
	if len(nums) == 1 {
		return 0
	}
	if mid < len(nums)-1 && abs(target-nums[mid]) > abs(target-nums[mid+1]) {
		return mid + 1
	}
	if mid >= 1 && abs(target-nums[mid]) > abs(target-nums[mid-1]) {
		return mid - 1
	}
	return mid
}
