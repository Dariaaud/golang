package main

import (
	"fmt"
)

func main() {
	//22 задача
	var si, el int
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	for i := 0; i < len(a); i += 1 {
		for j := 0; j < len(a)-1; j += 1 {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for i := 0; i < len(a); i += 1 {
		fmt.Println(a[i])
	}
}
