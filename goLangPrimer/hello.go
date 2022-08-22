package main

import (
	"fmt"
	"goLang/greetings"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	// message := greetings.Hello("Johny")
	// fmt.Println(message)
	// message, err := greetings.Hello("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)
	message1, err := greetings.Hello("Dipsy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message1)
}
