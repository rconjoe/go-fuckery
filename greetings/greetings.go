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

// Hellos returns a map that associates each of the names people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// make a map of key:value string:string
	messages := make(map[string]string)

	// in this for loop, range returns two values: the index and copy of the value.
	// we don't need the index so we can use blank identifier
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// assign a message to each name
		messages[name] = message
	}

	return messages, nil
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
