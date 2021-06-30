package main

import (
	"fmt"
	"strconv"
)

type Recepient interface {
	Set(i int)
}

type recepient struct {
	val int
}

func (r *recepient) Set(i int) {
	r.val = i
}

type Adapter struct {
	adaptee *recepient
}

type source struct {
	val string
}

func (s *source) Get() string {
	return s.val
}

func (a *Adapter) AdaptedSet(s string) {
	z, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	a.adaptee.Set(z)
}

func main() {
	var r recepient
	var s source
	s.val = "123"

	a := Adapter{
		&r,
	}
	a.AdaptedSet(s.Get())
	fmt.Println(r.val)
}
