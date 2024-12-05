package main

import (
	"sync"
)

func ExecuteTasksConcurrently(tasks []func(int, chan<- string, *sync.WaitGroup), numTasks int) ([]string, error) {
	var wg sync.WaitGroup
	results := make([]string, numTasks)
	ch := make(chan string, numTasks)

	for i, task := range tasks {
		wg.Add(1)
		go task(i, ch, &wg)
	}

	wg.Wait()
	close(ch)

	var i int
	for result := range ch {
		results[i] = result
		i++
	}

	return results, nil
}
