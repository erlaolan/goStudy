package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 5, 2, 6}
	countMap := findSingleNum(arr)
	for key, value := range countMap {
		if value == 1 {
			fmt.Println(key)
		}
	}
}

func findSingleNum(params []int) map[int]int {
	var countMap = make(map[int]int)
	for i := range params {
		value, exist := countMap[params[i]]
		if exist {
			countMap[params[i]] = value + 1
		} else {
			countMap[params[i]] = 1
		}
	}
	return countMap
}
