package main

import (
	"fmt"
	"math"
)

const half = 0.5

func square(side int) (area, perimeter int) {
	area = side * side
	perimeter = 4 * side
	return // Naked return
}

func rectangle(l, b int) (area, perimeter int) {
	fmt.Printf("RECTANGLE :::: LENGTH : %v | BREATH : %v | AREA: %v | PERIMETER : %v \n", l, b, l*b, 2*(l+b))
	fmt.Println("...............................................................................................................")
	return
}
func iso_triangle(h, b float64) (area, perimeter float64) {
	fmt.Printf("ISO_TRIANGLE :::: HEIGHT :%v | BREATH :%v | AREA: %v | PERIMETER : %v \n ", h, b, half*h*b, 2*h+b)
	fmt.Println("...............................................................................................................")
	return
}
func circle(r float64) (area, perimeter float64) {
	fmt.Printf("CIRCLE ::::RADUIS:%v | AREA: %v | PERIMETER : %v \n  ", r, math.Pi*r*r, 2*math.Pi*r)
	fmt.Println("...............................................................................................................")
	return
}

func scalene_traingle(a, b, c float64) (s, area, perimeter float64) {
	fmt.Printf("SCALENE_TRIANGLE :::: 1st side : %v | 2nd side : %v | 3rd side: %v | SEMI_PERIMETER : %v| AREA: %v | PERIMETER : %v \n ", a, b, c, (a+b+c)/2, math.Sqrt(s*(s-a)*(s-b)*(s-c)), a+b+c)
	fmt.Println("...............................................................................................................")
	return
}

//       3D figures

func cube(a int) {

	fmt.Printf("############ 3D SHAPES ############# \n \n \n ")
	fmt.Printf("CUBE:::: SIDE: %v | VOLUME : %v | TOTAL SURFACE AREA : %v \n", a, a*a*a, 6*a*a)
	fmt.Println("...............................................................................................................")
	return
}

func cuboid(l, b, h int) {
	fmt.Printf("CUBOID :::: LENGTH: %v | BREATH : %v | HEIGHT : %v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", l, b, h, l*b*h, 2*(l*b+b*h+h*b))
	fmt.Println("...............................................................................................................")
	return
}

func cylinder(r, h int) {

	fmt.Printf("CYLINDER :::: RADIUS: %v | HEIGHT :%v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", r, h, 22/7*r*r*h, 2*(22/7)*r*h+2*(22/7)*r*r)
	fmt.Println("...............................................................................................................")
	return
}

func cone(r, h, l int) {
	fmt.Printf("CONE :::: RADIUS: %v | HEIGHT :%v | LENGTH : %v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", r, h, l, (1/3)*(22/7)*r*r*h, (22/7)*r*(r+l))
	fmt.Println("...............................................................................................................")
	return
}

func main() {
	side := 6
	area, perimeter := square(side)
	fmt.Printf("Side: %v | Area: %v | Perimeter: %v\n", side, area, perimeter)
	fmt.Println("...............................................................................................................")

	rectangle(5, 6)
	circle(5)
	iso_triangle(4, 5)
	scalene_traingle(91, 55, 45)

	//3d shapes

	cube(4)
	cuboid(5, 6, 7)
	cylinder(4, 6)
	cone(4, 6, 8)

}
