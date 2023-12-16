package main

import "fmt"

func main() {
	var a, i, b rune
	a = '🂱'
	for i = 0; i < 13; i++{
		b = a + i
		fmt.Printf("Simbolo: %q Base 10: %d\n", b, b)
	}
}