package main

import "fmt"

func main() {
	var a, b, i, s int
	s = 0
	fmt.Scan(&a, &b)
	if a > b {
		i = b
		b = a
		a = i
	}
	i = a + 1
	for i < b {
		if i%2 == 0 { 
			s = s + i
		}
		i++
	}
	fmt.Println(s)
}