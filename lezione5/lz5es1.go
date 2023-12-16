package main

import "fmt"
func tabellina(a int) bool{
	var i, n int
	if a > 0 && a < 10 {
		fmt.Print("tabellina del ", a, ": " )
		for i = 0; i <= 10; i++{
			n = i * a
			fmt.Print(n," ")
		}
		return true
	} else {
		return false
	}
}
func main() {
	var a int 
	var i bool
	i = true
	for i == true {
		fmt.Scan(&a)
		i = tabellina(a)
	}
	if i == false{
		fmt.Println("Programma terminato")
	}
	}