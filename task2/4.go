package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskFunc func()
type Task struct {
	name string
	fn   TaskFunc
}

func main() {
	tasks := []Task{
		{
			name: "task1",
			fn: func() {
				time.Sleep(200 * time.Millisecond)
				fmt.Println("excute task1")
			},
		}, {
			name: "task2",
			fn: func() {
				time.Sleep(100 * time.Millisecond)
				fmt.Println("excute task2")
			},
		}, {
			name: "task3",
			fn: func() {
				time.Sleep(300 * time.Millisecond)
				fmt.Println("excute task3")
			},
		},
	}
	start := time.Now()
	wg := sync.WaitGroup{}

	for _, task := range tasks {
		wg.Add(1)

		go func(ta Task) {
			defer wg.Done()
			taskStart := time.Now()
			task.fn()
			taskDuration := time.Since(taskStart)
			fmt.Printf("%s excute time: %v\n", task.name, taskDuration)
		}(task)
	}
	wg.Wait()
	fmt.Printf("total excute time: %v\n", time.Since(start))
}
