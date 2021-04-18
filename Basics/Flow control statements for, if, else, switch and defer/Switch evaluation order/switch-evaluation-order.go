package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("Which month is ur Birthday?")
	today := time.Now().Month()
	switch time.January {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days")
	case today + 4:
		fmt.Println("In Four days.")
	case today + 5:
		fmt.Println("In Five days.")
	case today + 6:
		fmt.Println("In Six days.")
	case today + 7:
		fmt.Println("In Seven days.")
	case today + 8:
		fmt.Println("In Eight days")
	case today + 9:
		fmt.Println("In Nine days.")
	default:
		fmt.Println("Too far away.")
	}
}
