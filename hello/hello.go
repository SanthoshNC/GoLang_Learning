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
	
	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
	
	//message := greetings.Hello("Gladys")
	fmt.Println(message)
}
