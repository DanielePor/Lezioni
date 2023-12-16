package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inserisci testo (termina con CTRL+D):")

	testo := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}

	fmt.Print("Testo letto:\n", testo)

}