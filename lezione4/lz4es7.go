package main

import "fmt"

func main() {
	var b, i, j int
	fmt.Scan(&b)
	for i = 0; i < b; i++{
		fmt.Print("*")
	} 
	fmt.Print("\n")
	for j = 1; j < b ; j++{
		for i = 0; i<j; i++{
			fmt.Print(" ")
		}
		if j != b-1 {
		fmt.Print("*")
		}
		for i = 0; i< b-j-2; i++{
			fmt.Print(" ")
		}
		fmt.Println("*")	
	}
	for j = 0; j < b; j++{
		for i = 0; i < b-1; i++{
			fmt.Print(" ")
		} 
		if j != b-1 {
		fmt.Println("*")
		}
		}
	}
	
	