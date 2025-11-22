package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(arr))

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	arr := [][]int{intervals[0]}
	for _, val := range intervals[1:] {
		last := &arr[len(arr)-1]
		if (*last)[1] >= val[0] {
			if (*last)[1] < val[1] {
				(*last)[1] = val[1]
			}
		} else {
			arr = append(arr, val)
		}
	}

	return arr
}
