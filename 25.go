package main

import (
	"fmt"
	"sync"
)

type Counter interface {
	Set(i int)
	Inc() int
}

type counter struct {
	cnt int
	sync.RWMutex
	once sync.Once
}

var (
	cnt  *counter
	once sync.Once
)

func GetInstance() *counter {
	once.Do(func() {
		cnt = new(counter)
	})
	return cnt
}

func (s *counter) Set(i int) {
	s.Lock()
	s.cnt = i
	s.Unlock()
}

func (s *counter) Inc() int {
	s.Lock()
	s.cnt++
	defer s.Unlock()
	return s.cnt
}

func main() {
	cnt := GetInstance()
	wg := sync.WaitGroup{}
	cnt.Set(5)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		cnt := GetInstance()
		fmt.Println(cnt.Inc())
		wg.Done()
	}(&wg)
	wg.Wait()
}
