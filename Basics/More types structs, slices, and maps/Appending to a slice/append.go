package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3, 4)

	s = append(s, 10111, 11212)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d  cap=%d  %v \n", len(s), cap(s), s)

}
