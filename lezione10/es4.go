package main
import (
	"fmt"
	"os"
	"strings"
)
func main(){
	Sl:=os.Args[1:]
	Map:=Occorrenze(Sl)
	for i,v:=range Map{
		fmt.Printf("%s: %d\n",i,v)
	}
}
func Occorrenze(sl []string) map[string]int{
	var i,j int
	Mappa := make(map[string]int)
	var temp []string
	var temp1 string
	for i = range sl{
		for j = (i+1); j<len(sl);j++{
			temp=append(temp,sl[i:j+1]...)
			if temp[0] == temp[len(temp)-1] && len(temp)>=3{
				temp1=strings.Join(temp, " ")
				Mappa[temp1]+=1
			}
			temp1 = ""
			temp = []string{} 
		}
	}
	return Mappa
}