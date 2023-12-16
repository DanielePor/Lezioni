package main

import "fmt"
func Asterisco(l int){
	var i int
	for i = 0; i < l; i++{
		if i%2 == 0{
			fmt.Print("*")
		} else{
			fmt.Print("+")
		}
	}
	fmt.Println(" ")
	}
func Piu(l int){
	var i int
	for i = 0; i < l; i++{
		if i%2 == 0{
			fmt.Print("+")
		} else{
			fmt.Print("*")
		}
	}
	fmt.Println(" ")
	}
	
func Scacchiera(l int){
	var i int
	for i = 0; i < l; i++{
		if i%2 == 0{
			Asterisco(l)
		} else{
			Piu(l)
		}
	}
	}
func main() {
	var l int
	fmt.Scan(&l)
	Scacchiera(l)
	}