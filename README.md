# distance_between_2_array_go

This repo is for implementation for calculate distance between 2 array

## description

Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.

The distance value is defined as the number of elements arr1[i] such that there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.

## key point

first sort the array 2, then we could use binary search for the closed element for array 1 each element

time complexity:

1 sort array2 will be in O(NlogN) with heapsort

2 search array2 element will be O(logN) 

3 total search will be O(NlogN)

## implementation

```golang
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
```