package main

import "fmt"

func flipstr(s string) string {
	fl := []rune(s)
	for index := 0; index < len(fl)/2; index++ {
		fl[index], fl[len(fl)-index-1] = fl[len(fl)-index-1], fl[index]
	}
	return (string(fl))
}

func main() {
	fmt.Println(flipstr("HIPPITY HOPPITY WHERE THE FUCK IS MY PROPERTY"))
}
