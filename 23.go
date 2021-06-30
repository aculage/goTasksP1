package main

import "fmt"

func bsearch(a []int, offs int, val int) int {
	if a[len(a)/2] == val {
		return (len(a)/2 + offs)
	}
	if len(a)/2 == 0 {
		return -1
	}
	if a[len(a)/2] < val {
		return (bsearch(a[len(a)/2:], offs+len(a)/2, val))
	} else {
		return (bsearch(a[:len(a)/2], offs, val))
	}

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(bsearch(a, 0, 15))
}
