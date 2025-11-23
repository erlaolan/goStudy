package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
