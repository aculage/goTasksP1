package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan bool, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, stop chan bool) {
		for {
			select {
			case <-stop:
				fmt.Print("killed via direct channel send")
				wg.Done()
				return
			}
		}
	}(&wg, ch)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("killed via context expiration")
				return
			}
		}
	}(&wg, ctx)

	wg.Add(1)
	ctx1, cancel := context.WithCancel(context.Background())
	go func(wg *sync.WaitGroup, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("killed via context kill")
				return
			}
		}
	}(&wg, ctx1)

	var n int
	for {
		fmt.Scan(&n)
		if n == 5 {
			ch <- true
		}
		if n == 6 {
			cancel()
		}
		if n == 7 {
			break
		}
	}

	//TODO: Context Timeout and Context direct cancel call
	wg.Wait()
}
