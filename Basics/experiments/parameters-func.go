package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func square(a int) (area, perimeter int) {
	area = a * a
	fmt.Println("Area of square", area)
	perimeter = 4 * a
	fmt.Println("Perimeter of Square", perimeter)
	return
}

func rectangle(l, b int) (area, perimeter int) {
	area = l * b
	fmt.Println("Area of rectangle", area)
	perimeter = 2 * (l + b)
	fmt.Println("Perimeter of rectangle", perimeter)
	return
}
func iso_triangle(h, b float64) (area, perimeter float64) {
	area = 0.5 * h * b
	fmt.Println("Area of isos_Traingle", area)
	perimeter = 2*h + b
	fmt.Println("Perimeter of iso _triangle", perimeter)
	return
}
func circle(r float64) (area, perimeter float64) {
	area = Pi * r * r //constant
	fmt.Println("Area of circle", area)
	perimeter = 2 * Pi * r
	fmt.Println("Perimeter of circle", perimeter)
	return
}

func scalene_traingle(a, b, c float64) (s, area, perimeter float64) {
	s = (a + b + c) / 2
	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	fmt.Printf("Area of scalene_triangle", area)
	perimeter = a + b + c
	fmt.Printf("Perimeter of scalene_triangle", perimeter)
	return
}

func main() {
	fmt.Println(square(4))
	fmt.Println(rectangle(5, 6))
	fmt.Println(circle(5))
	fmt.Println(iso_triangle(4, 5))
	fmt.Println(scalene_traingle(4, 3, 3))

}
