package main

import "fmt"

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "A"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100]) //Protect slicing from cuttin Unicode
}

func main() {
	a := "нана"
	b := a[:3]
	fmt.Println(b)
	c := string([]rune(a)[:3])
	fmt.Println(c)
}
