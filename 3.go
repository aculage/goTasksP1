package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	ch <- 0
	for _, val := range a {
		wg.Add(1)
		go func(ch chan int, val int) {
			defer wg.Done()
			i := <-ch + val*val
			ch <- i
		}(ch, val)
	}
	wg.Wait()
	fmt.Println(<-ch)
}
