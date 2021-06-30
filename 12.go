package main

import "fmt"

func update(p *int /*p **int or this*/) {
	b := 2
	p = &b
	//*p = b either this, thus a is changed
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) //points to a
	update(p)       //gives a copy of pointer to a
	fmt.Println(*p) //thus points to a still
}
