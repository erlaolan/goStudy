package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2, 2)
	for idx, val := range nums {
		last := target - val
		idx2 := idx + 1
		for idx1, val1 := range nums[idx2:] {
			if last == val1 {
				res[0] = idx
				res[1] = idx + 1 + idx1
			}
		}
	}
	return res
}
