package main

import (
	"fmt"
)

func partition(a []int, lbound int, rbound int) int {
	p := a[rbound]

	for j := lbound; j < rbound; j++ {
		if a[j] < p {
			a[j], a[lbound] = a[lbound], a[j]
			lbound++
		}
	}
	a[lbound], a[rbound] = a[rbound], a[lbound]
	return lbound
}

func quickSort(a []int, lbound int, rbound int) {
	if lbound >= rbound {
		return
	}
	part := partition(a, lbound, rbound)
	quickSort(a, lbound, part-1)
	quickSort(a, part+1, rbound)
}

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
