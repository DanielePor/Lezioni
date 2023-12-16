package main

import "fmt"

func main() {
	var a, i, j, k int
	k = 0 
	fmt.Scan(&a)
	for  i = 0; i < a; i++  {
		if k == 0 {
			for j = 0; j < a; j++{
			fmt.Print("* ")
		}		
			k = 1
 		} else if k == 1{
			for j = 0; j < a; j++{
				fmt.Print("+ ")
			}
			k = 2	
		} else if k == 2 {
			for j = 0; j < a; j++{
				fmt.Print("o ")
			}
			k = 0
		} 
				
		fmt.Print("\n")
	}
}