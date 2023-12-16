package main

import "fmt"

func main() {
	var i, n, m int
	fmt.Scan(&n)
	for i = 1; i <= 10; i++ {
		m = n*i
		fmt.Println(n,"x",i,"=",m)
	} 

}