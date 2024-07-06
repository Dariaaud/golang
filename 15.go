package main

import (
	"fmt"
)

func mid(a []int) float32 {
	var ans float32
	for i := 0; i < len(a); i += 1 {
		ans += float32(a[i])
	}
	return ans / float32(len(a))
}

func main() {
	//15 задача
	var si, el int
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Println(mid(a))
}
