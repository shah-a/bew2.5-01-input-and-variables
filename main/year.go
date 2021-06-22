package main

import (
	"fmt"
	"log"
)

func year() {
	var age, dob int

	fmt.Print("How old will you be by the end of the year? -> ")
	if _, err := fmt.Fscanln(reader, &age); err != nil {
		log.Fatal(err)
	}

	fmt.Print("In what year were you born? -> ")
	if _, err := fmt.Fscanln(reader, &dob); err != nil {
		log.Fatal(err)
	}

	year := dob + age
	fmt.Println("\nThe current year is ->", year)
}
