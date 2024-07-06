package main

import (
	"fmt"
)

func elem(a []int, el int) bool {
	for i := 0; i < len(a); i += 1 {
		if a[i] == el {
			return true
		}
	}
	return false
}

func main() {
	//14 задача
	var si, el int
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&el)
	if elem(a, el) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
