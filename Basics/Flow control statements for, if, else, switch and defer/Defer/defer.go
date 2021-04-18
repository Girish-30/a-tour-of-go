package main

import "fmt"

func main() {
	defer fmt.Print("world ")
	defer fmt.Print("Duniya ")
	defer fmt.Print("Samaaj ")
	defer fmt.Print("Universe ")

	fmt.Printf("hello ")
}
