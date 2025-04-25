package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("asshole")
	fmt.Println(message)
}
