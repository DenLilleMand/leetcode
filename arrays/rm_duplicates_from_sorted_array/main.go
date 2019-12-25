package main

import (
	"fmt"
)

// Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
//
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

// Impl w/ extra space
func rm_dup(arr []int) int {
	tmp := []int{}
	for i, n := range arr {
		if i == 0 {
			tmp = append(tmp, n)
			continue
		}
		if n == arr[i-1] {
			continue
		}
		tmp = append(tmp, n)
	}
	return len(tmp)
}

// Impl w/ extra space
func rm_dup_using_map(sortedArr []int) int {
	tmp := map[int]interface{}{}
	for _, n := range sortedArr {
		tmp[n] = nil
	}
	r := []int{}
	for k, _ := range tmp {
		r = append(r, k)
	}
	return len(r)
}

// Accepted implementation
func rm_dup_using_const_space(sortedArr []int) int {
	length := len(sortedArr)
	r := 0
	for i := 0; i < length; i++ {
		remaining := false
		for j := i + 1; j < length; j++ {
			if sortedArr[i] >= sortedArr[j] {
				continue
			}
			remaining = true
			fmt.Printf("i: %+v, Arr:%+v\n", i, sortedArr)
			sortedArr[i+1] = sortedArr[j]
			break
		}
		if !remaining {
			return i + 1
		}
	}
	return r
}

// Accepted implementation
func rm_dup_using_const_space_optimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	sortedArr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r3 := rm_dup_using_const_space_optimized(sortedArr)
	fmt.Printf("len:%+v,arr:%+v\n", r3, sortedArr)
}
