package main
import (
	"fmt"
)
func main(){
	var x,y,temp int
	z:=1
	fmt.Scan(&x)
	for i:=0; i<x; i++{
		fmt.Printf("%d ",z)
		temp=z
		z += y
		y=temp
	}
}