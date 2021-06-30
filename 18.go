package main

import (
	"fmt"
	"unsafe"
)

func someAction(v []int8, b int8) {
	v[0] = 100 //access to internal structure
	fmt.Println(unsafe.Pointer(&v[0]))
	v = append(v, b) //local scope changes pointer
	fmt.Println(unsafe.Pointer(&v[0]))
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Pointer(&a[0]))
	someAction(a, 6)
	fmt.Println(unsafe.Pointer(&a[0]))
	fmt.Println(a)
}
