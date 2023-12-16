package main
import (
	"fmt"
	"os"
	"strconv"
	"sort"
	"strings"
)
func main(){
	var i, Num, y int
	var b, x1 []string
	var SliceNum []int
	x:=LeggiDividi()
	j:=3
	for y=0; y<3; y++{
	 	for i=0; i <(len(x)-(j-1)); i++{
			x1 = make([]string, len(x))
			copy(x1, x)
			b=append(x1[:i],x1[i+j:]...)
			Num = Converti(b)
		if Primo(Num) == true{
			SliceNum = append(SliceNum, Num)
		}
		b = nil
		}
		j--
	}
	sort.Ints(SliceNum)
	fmt.Println(SliceNum)
}
func LeggiDividi() (nsl []string){
	b:=os.Args[1:]
	nsl = strings.Split(b[0],"")
	return
}
func Converti(sl []string) (f int){
	num:=strings.Join(sl,"")
	f, _ = strconv.Atoi(num)
	return
}
func Primo(a int) bool{
	if a == 1 || a == 0{
		return false
	}
	for i:=2; i<a;i++{
		if a%i == 0{
			return false
		}
	}
	return true
}