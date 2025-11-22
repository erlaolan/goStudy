package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3}
	fmt.Println(x2(num))

}

func x2(target []int) []int {
	for idx, val := range target {
		target[idx] = 2 * val
	}
	return target
}
