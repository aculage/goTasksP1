package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var N int64

	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	fmt.Scan(&N)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)

	defer cancel()
	t := time.Now()

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%v\n", ctx.Err())
				wg.Done()
				return
			default:
				<-ch
			}
		}
	}(ctx, &wg, ch)

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Printf("%v\n", ctx.Err())
				wg.Done()
				return
			default:
				ch <- 5
			}
		}
	}(ctx, &wg, ch)

	wg.Wait()
	fmt.Println(time.Since(t))

}
