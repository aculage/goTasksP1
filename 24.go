package main

import "fmt"

func main() {
	a := make([]int, 100)
	fmt.Println(len(a), cap(a))
}
