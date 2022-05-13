package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with a message
	// In Go, you initialize a map with the following syntax:
	// make(map[key-type]value-type)
	messages := make(map[string]string)
	// Loop through the received slice of names, calling the
	// Hello function to get a message for each name.
	// In this for loop, range returns two values: the index of the
	// current item in the loop and a copy of the item's value.
	// You don't need the index, so you use the Go blank identifier
	// (an underscore) to ignore it.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name
		messages[name] = message
	}
	return messages, nil
}

// init set initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
