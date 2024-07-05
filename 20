package main

import (
	"fmt"
)

func lin(a []int, el int) int {
	for i := 0; i < len(a); i += 1 {
		if a[i] == el {
			return i + 1
		}
	}
	return -1
}

func main() {
	//20 задача
	var si, el int
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&el)
	fmt.Println(lin(a, el))
}
