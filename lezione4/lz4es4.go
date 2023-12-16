package main

import "fmt"

func main() {
	var a string
	var x,y,r byte
	x = 'A'
	y = 'Z'
	r = 'a'- 'A'
	X := x + r
	Y := y + r
	var i, M, m, v, c int
	fmt.Scan(&a)
	for i = 0; i < len(a); i++ {
		if a[i] >= x && a[i] <= y {
			M++
			if a[i] == x || a[i] == (x + 4) || a[i] == (x+8) || a[i] == (x + 14) || a[i] == (x + 20) {
				v++
			} else {
				c++
			}
		}
		if a[i] >= X && a[i] <= Y {
			m++
			if a[i] == X || a[i] == (X + 4) || a[i] == (X+8) || a[i] == (X + 14) || a[i] == (X + 20) {
				v++
			} else {
				c++
			}
		}
		

	}
	fmt.Println("Maiuscole:", M, "\nMinuscole:", m, "\nVolcali:", v, "\nConsonanti:", c)
	}
