package main

import (
	"fmt"
)

func lin(a string, b string, m map[byte]int) []byte {
	var ans []byte
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
	//26 задача
	m := make(map[byte]int)
	var a, b string
	fmt.Scan(&a, &b)
	var fin []byte
	fin = lin(a, b, m)
	if len(fin) == len(b) && len(fin) == len(a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
