package main

import (
	"fmt"
)

func main() {
	//23 задача
	var si, a1, a2 int
	a1 = 0
	a2 = 1
	fmt.Scan(&si)
	for i := 0; i < si; i += 1 {
		fmt.Println(a1)
		a1, a2 = a2, a1+a2
	}
}
