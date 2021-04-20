package main

import (
	"fmt"
)

func main() {

	a := 9
	var area int
	var perimeter int
	area = a * a
	perimeter = 4 * a
	fmt.Println("Area of square", area)
	fmt.Println("Perimeter of square", perimeter)
	fmt.Println(".......................................................................")

	l := 5
	b := 9
	var areaa int
	var perimeterr int
	areaa = l * b
	perimeterr = 2 * (l + b)
	fmt.Println("Area of rectangle", areaa)
	fmt.Println("Perimeter of rectangle", perimeterr)
	fmt.Println(".......................................................................")

	h := 7
	var areab int
	var perimeters int
	areab = 1 / 2 * h * b
	perimeters = 2*h + b
	fmt.Println("Area of isos_Traingle", areab)
	fmt.Println("Perimeter of iso _triangle", perimeters)
	fmt.Println(".......................................................................")

	r := 7
	var areaq int
	var perimeterq int
	areaq = 22 / 7 * r * r //constant
	perimeterq = 2 * 22 / 7 * r
	fmt.Println("Area of circle", areaq)
	fmt.Println("Perimeter of circle", perimeterq)
	fmt.Println(".......................................................................")

	m := 7
	n := 8
	q := 9
	var areaw int
	var perimeterw int
	s := (m + n + q) / 2
	areaw = (s * (s - m) * (s - n) * (s - q))
	perimeterw = m + n + q
	fmt.Println("Area of scalene_triangle", areaw)
	fmt.Println("Perimeter of scalene_triangle", perimeterw)
	fmt.Println(".......................................................................")
}
