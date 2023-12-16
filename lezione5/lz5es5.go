package main

import "fmt"
func Terna(a,b,c int) bool{
	var x,y,z int
	x = a * a
	y = b * b
	z = c * c
	if (x + y) == z{
		return true
	}
	return false
}
func NTerne(a int){
	var i,j,k int
	var x bool
	for i = 1; i < a; i++{
		for j = 1; j < a; j++{
			for k = 1; k < a; k++{
				x = Terna(i,j,k)
				if x == true{
					fmt.Print("(",i,",",j,",",k,") ")
				}
			}
		} 
	} 
}

func main() {
	var l int
	fmt.Scan(&l)
	if l >= 0{
		NTerne(l)
	}else{
		fmt.Println("la soglia non è positiva")
	}
}