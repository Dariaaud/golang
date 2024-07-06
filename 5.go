package main

import (
	"fmt"
)

func main() {
	//5 задача
	var a int
	fmt.Scan(&a)
	t := 1
	for i := 2; i <= a; i += 1 {
		t *= i
	}
	fmt.Println(t)
}
