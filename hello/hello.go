package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)


func main() {
	//fmt.Println("Hello, World!")
	
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }
	
	//message := greetings.Hello("Gladys")
	fmt.Println(message)
}
