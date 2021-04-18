package main

import "fmt"

func main() {
	values := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	fmt.Println(values)

	var a [3]string
	a[0] = " Duniya "
	a[1] = "mai aaya hai"
	a[2] = "to maar k jaa"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)
}
