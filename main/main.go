package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var exit bool

	for exit == false {
		var input string

		fmt.Println("===== MENU =====")
		fmt.Println("To run the year calculator, enter \"year\"")
		fmt.Println("To run the weight calculator, enter \"weight\"")
		fmt.Println("To exit, enter any other input.\n")
		fmt.Print("-> ")

		if _, err := fmt.Fscanln(reader, &input); err != nil {
			log.Fatal(err)
		}

		fmt.Println() // For terminal spacing

		switch input {
		case "year":
			year()
			fmt.Println() // For terminal spacing
		case "weight":
			weight()
			fmt.Println() // For terminal spacing
		default:
			exit = true
			fmt.Println("Bye-bye 🙂")
			fmt.Println() // For terminal spacing
		}
	}
}
