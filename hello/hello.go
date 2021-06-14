package main

import "fmt"

//Adding to call the external package

import "rsc.io/quote"

func main() {
	//fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
