package main

import (
	"fmt"
	"reflect"
)

func gettype_reflect(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func gettype_assertion(i interface{}) {
	switch i.(type) {
	case int, int32, int64:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case chan int:
		fmt.Println("channel")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	ch := make(chan int)
	gettype_reflect("a")
	gettype_reflect(1)
	gettype_reflect([]int{1, 2, 3})
	gettype_reflect(ch)

	gettype_assertion("a")
	gettype_assertion(1)
	gettype_assertion([]int{1, 2, 3})
	gettype_assertion(ch)
}
