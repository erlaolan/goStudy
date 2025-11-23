package main

import (
	"fmt"
	"sync"
)

func sendOnly(ch chan<- int) {
	for i := 1; i < 101; i++ {
		ch <- i
		fmt.Printf("send---%v\n", i)
	}
	close(ch)
}

func receiveOnly(ch <-chan int) {
	for val := range ch {
		fmt.Printf("receive---%v\n", val)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	wg.Add(2)

	go func() {
		defer wg.Done()
		sendOnly(ch)
	}()

	go func() {
		defer wg.Done()
		receiveOnly(ch)
	}()

	wg.Wait()

}
