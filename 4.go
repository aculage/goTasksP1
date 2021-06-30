package main

import (
	"context"
	"fmt"
	"sync"
)

type Worker struct {
}

func (w *Worker) Run(ctx context.Context, ch chan interface{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker died")
				return
			default:
				switch z := <-ch; z {
				case nil:
					fmt.Println("worker natural shutdown")
					return
				default:
					fmt.Printf("%v ", z)
				}

			}
		}
	}()
}

func main() {
	var N int = 0
	fmt.Scan(&N)
	ctx, cancel := context.WithCancel(context.Background())
	workers := make([]Worker, N)
	wg := sync.WaitGroup{}
	ch := make(chan interface{}, 2*N)
	for _, w := range workers {
		w.Run(ctx, ch, &wg)
	}
	sendersWg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		sendersWg.Add(1)
		go func(val int, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				ch <- val
			}
		}(i, &sendersWg)
		/*if i == 30 {
			fmt.Println("killall attempt")
			cancel()
			fmt.Println(<-ctx.Done())
		}*/
	}
	sendersWg.Wait()
	close(ch)
	cancel()
	wg.Wait()
}
