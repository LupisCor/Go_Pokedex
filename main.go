package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		data := scanner.Text()
		cleanData := cleanInput(data)
		fmt.Printf("Your command was: %v \n", cleanData[0])
	}
}
