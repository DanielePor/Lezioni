package main
import  (
		"fmt"
		"os"
 		"strconv"
)
func main(){
	var s int
	a:=LeggiNumeri()
	fmt.Println(a)
	for _, x := range a{
		b:=Occorrenze(a,x)
		if b == 1{
			s = s+x
		}
	}
	fmt.Println(s)
}
func LeggiNumeri() []int{
	sl:=os.Args[1:]
	var f []int
	for _, s:= range sl{
		n, e := strconv.Atoi(s)
		if e == nil{
			f = append (f,n)
		}
	}
	return f
}
func Occorrenze(numeri []int, n int) int{
	var c int
	for _, x:= range numeri{
		if n==x{
			c++
		}
	}
	return c
}