package main

import "fmt"

func Unique(s string) bool {
	m := make(map[rune]bool)
	rs := []rune(s)

	for _, sym := range rs {
		if m[sym] {
			return false
		}
		m[sym] = true
	}
	return true
}

func main() {

	s := "QWERTYUIOPP" // check for lcase upcase?

	fmt.Println(Unique(s))
}
