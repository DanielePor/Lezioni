package main
import (
	"fmt"
	"os"
	"strconv"
)
func main(){
	a:=LeggiNumeri()
	fmt.Println(a)
	fmt.Println(Calcola(a))
}
func Calcola(sl []int) int{
	i:=0
	s:=0
	for i < len(sl){
		if i%2 == 0{
			s = s + (sl[i]*sl[i+1])
			i++
		} else{
			i++
		}
	}
	return s
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