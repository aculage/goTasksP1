package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	i := 5
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
}
