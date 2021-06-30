package main

import "fmt"

func main() {
	a := 100
	b := 2
	a, b = b, a
	fmt.Println(a, b)
}
