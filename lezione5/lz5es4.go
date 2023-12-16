package main

import "fmt"
func Primo(a int) bool{
	var i int
	for i = 2; i < a; i++{ 
		if a%i == 0{
			return false
		}
	}
	return true
}
func NGemelli(a int){
	var i,j int
	var x,y bool
	for i = 2; i < a-2; i++{
		x = Primo(i)
		j = i + 2
		y = Primo(j)
		if x == true && y == true{
			fmt.Print("(",i,",",j,")")
		} 
	}
}
func main() {
	var l int
	fmt.Scan(&l)
	if l >= 0{
		NGemelli(l)
	}else{
		fmt.Println("la soglia non è positiva")
	}
}