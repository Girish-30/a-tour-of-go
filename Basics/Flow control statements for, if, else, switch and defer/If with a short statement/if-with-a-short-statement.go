package main

import (
	"fmt"
)

func add(a, b, c float64) float64 {
	if d := a - b - c; d > 0 {
		return d
	}
	return c
}
func main() {
	fmt.Println(5, 10, -90)

}
