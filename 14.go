package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool)
	for _, v := range strings {
		m[v] = true
	}
	defer fmt.Println()
	for key, _ := range m {
		fmt.Printf("%v ", key)
	}
}
