package main

import "fmt"

func main() {
	var n, a, b, c int
	b = 1
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println(b)
		c = a
		a = b
		b = a + c
	}

}