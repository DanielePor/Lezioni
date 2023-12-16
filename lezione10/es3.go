package main
import (
	"fmt"
	"bufio"
	"os"
	"unicode"
)
func main(){
	text:=LeggiTesto()
	fmt.Println(text)
	SlText:= SeparaParole(text)
	fmt.Println(SlText)
	Mappa:=ContaRipetizioni(SlText)
	for k,v:= range Mappa{
		fmt.Printf("%s: %d\n", k, v)
	}
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
func SeparaParole(s string) []string{
	SlRune:=[]rune(s)
	var temp []rune
	var SlString []string
	for _,v := range SlRune{
		if unicode.IsLetter(v){
			temp=append(temp,v)
		}else{
			if len(temp) != 0{
			SlString=append(SlString,string(temp))
			}
			temp = []rune{}
		}
	}
	if len(temp) != 0{
		SlString=append(SlString,string(temp))
	}
	fmt.Printf("Parole contate: %d\n",len(SlString))
	return SlString
}
func ContaRipetizioni(sl []string) map[string]int{
	Map:=make(map[string]int)
	for _,k:= range sl{
		Map[k] += 1
	}
	return Map
}