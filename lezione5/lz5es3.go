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
func NPrimi(a int){
	var i int
	var j bool
	for i = 2; i < a; i++{
		j = Primo(i)
		if j == true{
			fmt.Print(i," ")
		} 
	}
}

func main() {
	var l int
	fmt.Scan(&l)
	if l >= 0{
		NPrimi(l)
	}else{
		fmt.Println("la soglia non è positiva")
	}
}