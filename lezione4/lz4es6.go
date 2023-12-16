package main

import "fmt"

func main() {
	var s string
	var f, i, x int
	f = 1
	fmt.Scan(&s)
	x = len(s) - 1
	for i = 0; i < len(s); i++{
		if s[i] != s [x-i] {
			f = 0
			fmt.Println("non palindroma")
			break
		}
	} 
	if f == 1{
		fmt.Println("palindroma")
	}
	}