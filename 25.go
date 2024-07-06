package main

import (
	"fmt"
)

func lin(a []int, b []int, m map[int]int) []int {
	var ans []int
	for i := 0; i < len(a); i += 1 {
		_, c := m[a[i]]
		if !c {
			m[a[i]] = 0
		}
		m[a[i]] += 1
	}
	for i := 0; i < len(b); i += 1 {
		_, c := m[b[i]]
		if !c || m[b[i]] == 0 {
			continue
		}
		m[b[i]] -= 1
		ans = append(ans, b[i])
	}
	return ans
}

func main() {
	//25 задача
	var si, si1, el int
	m := make(map[int]int)
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&si1)
	var b []int
	for i := 0; i < si1; i += 1 {
		fmt.Scan(&el)
		b = append(b, el)
	}
	a = lin(a, b, m)
	fmt.Println(a)
}
