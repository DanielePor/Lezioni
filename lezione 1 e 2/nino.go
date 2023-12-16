package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a)
	b = a % 100
	a = a / 100
	fmt.Print(b,a)
}