package main

import "fmt"

func main() { //scope0
	n := 0
	if true { //scope 1
		n := 1
		n++
	} //scope 1

	fmt.Println(n)
} //scope 0
