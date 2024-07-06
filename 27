package main

import (
	"fmt"
)

func merge(a []int, b []int) []int {
	var ans []int
	now1 := 0
	now2 := 0
	for now1+now2 < len(a)+len(b) {
		if now1 == len(a) || (now2 != len(b) && a[now1] > b[now2]) {
			ans = append(ans, b[now2])
			now2 += 1
		} else {
			ans = append(ans, a[now1])
			now1 += 1
		}
	}
	return ans
}

func main() {
	//27 задача
	var si, si1, el int
	var a, b []int
	fmt.Scan(&si)
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&si1)
	for i := 0; i < si1; i += 1 {
		fmt.Scan(&el)
		b = append(b, el)
	}
	a = merge(a, b)
	fmt.Println(a)
}
