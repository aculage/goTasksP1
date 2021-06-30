package main

import (
	"fmt"
	"math"
)

type Point interface {
	Create(x, y float64) *point
}

type point struct {
	x, y float64
}

func Create(x, y float64) *point {
	p := new(point)
	p.x = x
	p.y = y
	return p
}

func (p *point) Distance(p2 *point) float64 {
	return math.Sqrt((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y))
}

func main() {
	p1 := Create(1, 2)
	p2 := Create(5, 6)
	fmt.Println(p1.Distance(p2))
}
