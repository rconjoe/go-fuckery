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

	// make slice of names
	names := []string{"trogdor", "rconjoe", "joe"}

	// run Hellos over the names slice
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
