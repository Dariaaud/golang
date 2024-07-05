package main

import (
	"fmt"
)

var ans1, ans2 int

func mima(a []int) {
	for i := 0; i < len(a); i += 1 {
		ans1 = min(ans1, a[i])
		ans2 = max(ans2, a[i])
	}
}

func main() {
	//18 задача
	var si, el int
	ans1 = 1000000000
	ans2 = -1
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	mima(a)
	fmt.Println(ans1, ans2)
}
