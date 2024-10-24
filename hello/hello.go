package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())
	greet("mengxiangjian")
	greet("guozhenyan")
	greets([]string{"mengxiangjian", "guozhenyan"})
	greet("")
}

func greet(name string) {
	message, error := greetings.Hello(name)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}

func greets(names []string) {
	messages, error := greetings.Hellos(names)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(messages)
}
