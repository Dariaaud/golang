package main

import (
	"fmt"
)

func sli(a []int, el int) []int {
	var ans []int
	ans = a[:el-1]
	for i := el; i < len(a); i += 1 {
		ans = append(ans, a[i])
	}
	return ans
}

func main() {
	//19 задача
	var si, el int
	fmt.Scan(&si)
	var a []int
	for i := 0; i < si; i += 1 {
		fmt.Scan(&el)
		a = append(a, el)
	}
	fmt.Scan(&el)
	a = sli(a, el)
	for i := 0; i < len(a); i += 1 {
		fmt.Println(a[i])
	}
}
