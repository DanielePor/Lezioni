package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for a > -1 {
		b = a
		fmt.Scan(&a)
		if a < 0{
			fmt.Print("Sequenza finita")
			break
		}
		if a >= b{
			fmt.Println("crescente")
		} else {
			fmt.Println("decrescente")
		}
	}
}