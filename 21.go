package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg := sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, a []int) {
			defer fmt.Println()
			for _, v := range a {
				fmt.Printf("%v ", v)
			}
			wg.Done()
		}(&wg, arr[i*(len(arr)/4):(i+1)*(len(arr)/4)])
	}
	wg.Wait()
}
