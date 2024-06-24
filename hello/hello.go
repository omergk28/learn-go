package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Ali", "Veli", "Mustafa"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
