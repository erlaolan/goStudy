package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			if i%2 == 1 {
				fmt.Println("odd---%d", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			if i%2 == 0 {
				fmt.Println("even---%d", i)
			}
		}
	}()
	wg.Wait()
}

func x2(target []int) []int {
	for idx, val := range target {
		target[idx] = 2 * val
	}
	return target
}
