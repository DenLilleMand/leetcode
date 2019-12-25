package main

import (
	"fmt"
	"sort"
)

// Just say if it contains dups or not, the fastest
// solution in Go is to actually just sort the array and compare all
// nums[i] with nums[i-1] for equality

func containsDuplicate(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	s := map[int]interface{}{}
	for _, v := range nums {
		if _, ok := s[v]; ok {
			return true
		} else {
			s[v] = nil
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	input := []int{1, 2, 3, 1}
	input2 := []int{2, 2}
	fmt.Println(containsDuplicate2(input))
	fmt.Println(containsDuplicate2(input2))

}
