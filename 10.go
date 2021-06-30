package main

import "fmt"

func main() {
	disp := make(map[int64][]float32)
	temps := []float32{-25.4, -27.0, 13.3, 19.0, 15.5, 24.5, -21.0, 32.5}
	if len(temps) == 0 {
		return
	}
	for _, v := range temps {
		disp[int64(v)/10*10] = append(disp[int64(v)/10*10], v)
	}
	fmt.Println(disp)
}
