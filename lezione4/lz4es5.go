package main

import "fmt"

func main() {
	var a, b, c string
	var i int
	var x, y byte
	y = 'a' - 'A'
	fmt.Scan(&a)
	for i = 0; i < len(a); i++ {
		if a[i] >= 'a' && a[i]<= 'z'{
			x = a[i] - y
			b += string(x)
		} else {
			x = a[i]
			b += string(x)
		}
	}
	for i = 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i]<= 'Z'{
			x = b[i] + y
			c += string(x)
		} else{
			c += string(b[i])
 		}
	}
	fmt.Println(string(b))
	fmt.Println(string(c))
	}