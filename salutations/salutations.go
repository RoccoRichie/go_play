package main

import (
	"fmt"

	"github.com/RoccoRichie/go_play/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Richie")
	fmt.Println(message)
}