package main

import "fmt"

func main() {
	fmt.Println("Benvenuto")
	var i, c, p int
	fmt.Print("1. Ritiro\n", "2. Consegna (2€)\n")
	fmt.Scan(&c)
	if c == 2 {
		p = 2
	}
	for c!=0 {
		fmt.Println("Cosa vuole ordinare?")
		fmt.Println("1. Patatine 2€\n2. Hamburger 5€\n3. CocaCola 3€\n0. Termina")
		fmt.Scan(&c)
		switch c {
			case 0: break
			case 1: {
				fmt.Println("Patatine? Ottimo quante?")
				fmt.Scan(&i)
				p = p + 2 * i 
			}
			case 2:{
				fmt.Println("Hamburger? Ottimo quanti?")
				fmt.Scan(&i)
				p = p + 5 * i 
			}
			case 3:{
				fmt.Println("CocaCola? Ottimo quante?")
				fmt.Scan(&i)
				p = p + 3 * i 
			}
		}
	}
	fmt.Print("Ottimo, il prezzo totale è di ", p, "€\n")
}