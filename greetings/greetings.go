package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages.
func randomFormat() string {
	// a slice of message formats
	// When declaring a slice, you omit its size in the brackets, like this: []string.
	// This tells Go that the size of the array underlying the slice can be dynamically changed.
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see, you %v",
		"What's up %v",
	}

	// return a format from a random index
	return formats[rand.Intn(len(formats))]
}
