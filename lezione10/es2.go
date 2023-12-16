package main
import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"sort"
	"strings"
)
func main(){
	Text:=LeggiTesto()
	MappaOcc:=Occorrenze(Text)
	StampaIstogramma(MappaOcc)
}
func Occorrenze(s string) map[rune]int{
	Mappa:=make(map[rune]int)
	SlRune:= []rune(s)
	for _,k:=range SlRune{
		if unicode.IsLetter(k){
			Mappa[k]+=1
		}
	}
	return Mappa
}
func LeggiTesto() string{
	fmt.Println("Inserisci testo (termina con CTRL+D):")
	testo := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return testo
}
func StampaAsterischi(x int){
	for i:=0;i<x;i++{
		fmt.Print("*")
	}
}
func StampaIstogramma(occorrenze map[rune]int){
	SlRune:=[]rune{}
	for k := range occorrenze {
		SlRune=append(SlRune,k)
	}
	SlString:=string(SlRune)
	CIAO:=strings.Split(SlString,"")
	sort.Strings(CIAO)
	SUCA:=strings.Join(CIAO,"")
	BOH:=[]rune(SUCA)
	for _,v:= range BOH{
		fmt.Printf("%c",v)
		StampaAsterischi(occorrenze[v])
		fmt.Print("\n")
	}
}