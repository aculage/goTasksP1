package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func WgOnly() {
	a := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, val := range a {
		wg.Add(1)
		go func(val int) {
			fmt.Println(val * val)
			wg.Done()
		}(val)
	}
	wg.Wait()
}

func WgAndChan() {
	a := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	for _, val := range a {
		wg.Add(1)
		go func(val int) {
			ch <- (val * val)
			wg.Done()
		}(val)
	}
	wg.Wait()
	close(ch)
	for z := range ch {
		fmt.Println(z)
	}
}

func ChanOnly() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5)
	for _, val := range a {
		go func(val int) {
			ch <- (val * val)
		}(val)
	}
	close(ch)
	for z := range ch {
		fmt.Println(z)
	}
}

func TestAvgTime(count int64, f func(), funcname string, bfw *os.File) {
	var z time.Duration
	for i := 0; int64(i) < count; i++ {
		a := time.Now()
		f()
		z += time.Since(a)
	}
	fmt.Fprintf(bfw, "Average execution time for %v is %v\n", funcname, time.Duration(int64(z)/count))
}
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	TestAvgTime(50, WgAndChan, "Wait Group and Channel", file)
	TestAvgTime(50, WgOnly, "Wait Group Only", file)
	TestAvgTime(50, ChanOnly, "Chan Only", file)
	file.Close()
}
