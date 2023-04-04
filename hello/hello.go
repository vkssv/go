package main

import (
	"fmt"
	"log"

	"example/greetings"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(1)

	names := []string{"VK", "K", "A"}

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())

	message, err := greetings.Hello("V")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	messages, e := greetings.Hellos(names)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(messages)
}
