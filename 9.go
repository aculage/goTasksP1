package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := make(chan int)
	dbls := make(chan int)
	wg := sync.WaitGroup{}
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg.Add(1)
	go func(wg *sync.WaitGroup, a []int, nums chan int) {
		for _, val := range a {
			nums <- val
		}
		close(nums)
		wg.Done()
	}(&wg, a, nums)
	wg.Add(1)
	go func(wg *sync.WaitGroup, dbls chan int, nums chan int) {
		for val := range nums {
			dbls <- 2 * val
		}
		close(dbls)
		wg.Done()
	}(&wg, dbls, nums)
	wg.Add(1)
	go func(wg *sync.WaitGroup, dbls chan int, nums chan int) {
		for val := range dbls {
			fmt.Printf("%v ", val)
		}
		wg.Done()
	}(&wg, dbls, nums)
	wg.Wait()
	fmt.Println()
}
