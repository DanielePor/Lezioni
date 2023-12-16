package main
import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)
func main(){
	Text:=LeggiTesto()
	MappaOcc:=Occorrenze(Text)
	for k,v := range MappaOcc {
		fmt.Printf("\n%c ",k)
		StampaAsterischi(v)
	}
	
}
func Occorrenze(s string) map[rune]int{
	Mappa:=make(map[rune]int)
	SlRune:= []rune(s)
	fmt.Printf("%U",SlRune)
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