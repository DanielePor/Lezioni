package main
import (
	"fmt"
	"os"
	"strings"
)
func main(){
	b:=os.Args[1:]
	Parola:=b[0]
	strings.ToLower(Parola)
	Occorrenze(Parola)
}
func Occorrenze (s string){
	var Occ [26]int
	Alpha:=[]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for i, v := range Alpha{
		for j:= 0;j<len(s);j++ {
			if string(s[j]) == v{
				Occ[i]++
			}
		}
	}
	for i:=0;i<26;i++{
		fmt.Printf("%v %v ",Occ[i],Alpha[i])
	}
}