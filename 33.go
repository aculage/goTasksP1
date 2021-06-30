package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch1 chan int) {
		rand.Seed(time.Now().UnixNano())
		defer close(ch1)
		for i := 0; i < 50; i++ {
			ch1 <- rand.Int() % 1000
		}
		wg.Done()
	}(&wg, ch1)

	wg.Add(1)
	go func(wg *sync.WaitGroup, ch1, ch2 chan int) {
		defer close(ch2)
		for z := range ch1 {
			if z%2 == 0 {
				ch2 <- z
			}
		}
		wg.Done()
	}(&wg, ch1, ch2)

	wg.Add(1)
	go func(wg *sync.WaitGroup, ch2 chan int) {
		for z := range ch2 {
			fmt.Printf("%v ", z)
		}
		wg.Done()
	}(&wg, ch2)

	wg.Wait()
	fmt.Println()
}
