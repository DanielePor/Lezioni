package main
import (
	"fmt"
	"math/rand"
	"time"
	"errors"
)
type Carta struct {
	seme string
	simbolo string
}

type Mazzo struct{
	Carte []Carta
	nCarte int
}

func main() {
	var Mazzo1 Mazzo
	Mazzo1 = CreaMazzo()
	var Carta1 Carta
	var e error
	Mischia(Mazzo1)
	fmt.Println(Mazzo1.Carte,Mazzo1.nCarte)
	Carta1,Mazzo1,e = Preleva(Mazzo1)
	fmt.Println(Carta1,Mazzo1.Carte,Mazzo1.nCarte, e)
	Mazzo1,e=Riponi(Mazzo1,Carta1)
	fmt.Println(Mazzo1.Carte,Mazzo1.nCarte)
}

func CreaCarta(seme,simbolo string) Carta{
	var x Carta
	x.seme=seme
	x.simbolo=simbolo
	return	x
}

func CreaMazzo() Mazzo{
	var semi = []string{"quadri", "picche", "cuori", "fiori"}
	var simboli = []string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette","otto","nove","dieci", "fante","donna", "re"}	
	var x Mazzo
	var y Carta
	x.nCarte=(len(semi)*len(simboli))
	for i:=0;i<len(semi); i++{
		for j:=0;j<len(simboli);j++{
			y = CreaCarta(semi[i],simboli[j])
			x.Carte=append(x.Carte, y)
		}
	}
	return x
}

func Mischia(m Mazzo){
	rand.Seed(int64(time.Now().Nanosecond())) 
	for i := range m.Carte {
 	j := rand.Intn(i + 1)
 	m.Carte[i], m.Carte[j] = m.Carte[j], m.Carte[i]
 }
	return
}

func Preleva(m Mazzo) (Carta, Mazzo, error){
	var c Carta
	var err error 
	if len(m.Carte) > 0 {
		c = m.Carte[0]
		m.Carte = m.Carte[1:]
		m.nCarte--
		err=nil
	}else{
		err=errors.New("Mazzo vuoto")
	}
	return c, m, err
}

func Riponi(m Mazzo, c Carta) (Mazzo, error){
	var err error
	if len(m.Carte)>0{
		m.Carte = append([]Carta{c} , m.Carte...)
		m.nCarte = len(m.Carte)
		err=nil
	}else{
		err=errors.New("Mazzo vuoto")
	}
	return m, err
}