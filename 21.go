package main

import (
	"fmt"
)

func lin(a []int, m map[int]int) []int {
	var ans []int
	for i := 0; i < len(a); i += 1 {
		_, b := m[a[i]]
		if b {
			continue
		}
		m[a[i]] = 1
		ans = append(ans, a[i])
	}
	return ans
}

func main() {
	//21 задача
	var si, el int
	m := make(map[int]int)
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	a = lin(a, m)
	for i := 0; i < len(a); i += 1 {
		fmt.Println(a[i])
	}
}
