package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	// set log entry prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("asshole")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
