package main

import (
	"fmt"
	"math"
)

const Pi = 3.14
const half = 0.5

func square(a int) (area, perimeter int) {
	area = a * a
	fmt.Println("Area of square", area)
	perimeter = 4 * a
	fmt.Println("Perimeter of Square", perimeter)
	fmt.Println(".......................................................................")
	return
}

func rectangle(l, b int) (area, perimeter int) {
	area = l * b
	fmt.Println("Area of rectangle", area)
	perimeter = 2 * (l + b)
	fmt.Println("Perimeter of rectangle", perimeter)
	fmt.Println(".......................................................................")
	return
}
func iso_triangle(h, b float64) (area, perimeter float64) {
	area = half * h * b
	fmt.Println("Area of isos_Traingle", area)
	perimeter = 2*h + b
	fmt.Println("Perimeter of iso _triangle", perimeter)
	fmt.Println(".......................................................................")
	return
}
func circle(r float64) (area, perimeter float64) {
	area = Pi * r * r //constant
	fmt.Println("Area of circle", area)
	perimeter = 2 * Pi * r
	fmt.Println("Perimeter of circle", perimeter)
	fmt.Println(".......................................................................")
	return
}

func scalene_traingle(a, b, c float64) (s, area, perimeter float64) {
	s = (a + b + c) / 2
	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	fmt.Println("Area of scalene_triangle", area)
	perimeter = a + b + c
	fmt.Println("Perimeter of scalene_triangle", perimeter)
	fmt.Println(".......................................................................")
	return
}

func main() {
	square(4)
	rectangle(5, 6)
	circle(5)
	iso_triangle(4, 5)
	scalene_traingle(4, 3, 3)
}
