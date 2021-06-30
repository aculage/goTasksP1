package main

import "sync"

var lock sync.RWMutex

func read(m map[interface{}]interface{}, key interface{}) interface{} {
	lock.RLock()
	defer lock.RUnlock()
	return m[key]
}

func write(m map[interface{}]interface{}, key interface{}, val interface{}) {
	lock.Lock()
	defer lock.Unlock()
	m[key] = val
}

func main() {
	a := make(map[int]int)
}
