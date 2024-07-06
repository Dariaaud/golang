package main

import (
	"fmt"
)

func main() {
	//29 задача
	var si, el int
	var a []int
	fmt.Scan(&si)
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&el)
	var m int
	l := 0
	r := si
	for r-l > 1 {
		m = (l + r) / 2
		if a[m] > el {
			r = m
		} else {
			l = m
		}
	}
	fmt.Println(l + 1)
}
