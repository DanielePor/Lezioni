package main

import "fmt"

func main() {
	var a,b,c,d , x, y float64
	fmt.Scan(&a,&b,&c,&d)
	fmt.Println(a,"/",b,"+",c,"/",d)
	var numero float64 = (a/b)+(c/d)
	fmt.Println("risultato in numero",numero)
	a = a*d
	c = c*b
	y = b*d
	x = a+c
	fmt.Println("risultato in frazione",x,"/",y)

}