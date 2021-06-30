package main

import "fmt"

type Action struct {
	val interface{}
}

type Human struct {
	action Action
}

func (this *Human) Init(a string) {
	this.action = Action{
		a,
	}
}

func main() {
	var hooman Human
	hooman.Init("Initialized")
	hooman.action = Action{}
	fmt.Println(hooman.action.val)
}
