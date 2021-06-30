package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "HIPPITY HOPPITY WHERE THE FUCK IS MY PROPERTY"
	words := strings.Split(str, " ")
	newstr := ""
	for i := len(words) - 1; i >= 0; i-- {
		newstr += words[i]
		if i != 0 {
			newstr += " "
		}
	}
	fmt.Println(newstr)
}
