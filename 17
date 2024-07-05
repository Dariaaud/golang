package main

import (
	"fmt"
)

func poly(s string) bool {
	for i := 0; i <= len(s)/2; i += 1 {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//17 задача
	var s string
	fmt.Scan(&s)
	if poly(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
