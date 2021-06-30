package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {

	timer := time.NewTimer(t)
	<-timer.C
	return
}

func main() {
	t := time.Now()
	Sleep(5 * time.Second)
	fmt.Println(time.Since(t))
}
