package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "a"}

	//fmt.Println(unsafe.Pointer(&slice[0]))
	func(slice []string) {
		slice = append(slice, "a") // whoops we r remapping
		//fmt.Println(unsafe.Pointer(&slice[0]))
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice) // b b a
	}(slice)
	//fmt.Println(unsafe.Pointer(&slice[0]))
	fmt.Print(slice) // a a
}
