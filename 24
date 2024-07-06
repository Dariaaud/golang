package main

import (
	"fmt"
)

func count(a []int, el int) int {
	ans := 0
	for i := 0; i < len(a); i += 1 {
		if a[i] == el {
			ans += 1
		}
	}
	return ans
}

func main() {
	//24 задача
	var si, el int
	var a []int
	fmt.Scan(&si)
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&el)
	fmt.Println(count(a, el))
}
