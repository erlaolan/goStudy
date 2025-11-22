package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	i := 0
	if len(nums) == 0 {
		return i
	}
	for _, val := range nums {
		if nums[i] != val {
			i++
			nums[i] = val
		}
		fmt.Println(nums)
	}
	return i + 1
}
