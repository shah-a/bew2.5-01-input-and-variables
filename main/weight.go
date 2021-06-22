package main

import (
	"fmt"
	"log"
)

func weight() {
	var w1, w2, w3, w4, w5 float64

	fmt.Println("Enter the weights of 5 people (in one line, separated by spaces)")
	fmt.Print("-> ")
	if _, err := fmt.Fscanln(reader, &w1, &w2, &w3, &w4, &w5); err != nil {
		log.Fatal(err)
	}

	avgWeight := (w1 + w2 + w3 + w4 + w5) / 5
	fmt.Printf("\nThe average weight is -> %.2f\n", avgWeight)
}
