package main

import (
	"fmt"

	"github.com/NYTgophers/casing-go/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gophers")
	fmt.Println(message)
}
