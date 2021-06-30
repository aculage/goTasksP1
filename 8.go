package main

import (
	"fmt"
)

func setBit(pos int, val int64, bit bool) int64 {
	if bit {
		return (val | (int64(1) << (pos - 1)))
	} else {
		return (val & ^(int64(1) << (pos - 1)))
	}
}

func main() {

	var a int64 = 0

	fmt.Println(a)
	fmt.Println(setBit(4, a, true)) //8 expected
	a = 4
	fmt.Println(setBit(3, a, false)) // 0 expected

}
