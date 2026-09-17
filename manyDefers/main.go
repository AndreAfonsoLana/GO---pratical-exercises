package main

import "fmt"

func main() {
	fmt.Println("Welcome the ")
	TesteDefer()
}

func TesteDefer() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")

	fmt.Println("Codigo meio")
}
