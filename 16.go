package main

import (
	"fmt"
)

func main() {
	//16 задача
	var el int
	fmt.Scan(&el)
	for i := 1; i <= 9; i += 1 {
		fmt.Println(el, "*", i, "=", el*i)
	}
}
