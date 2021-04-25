package main

import (
	"fmt"
)

const Pi = 22 / 7
const half = 1 / 2

func main() {

	fmt.Println(" \n ############################################## 2D SHAPE ######################################### \n ")
	fmt.Println("...............................................................................................................")

	side := 9
	fmt.Printf("SQUARE ::::: Side: %v | Area: %v | Perimeter: %v \n", side, side*side, 4*side)
	fmt.Println(".......................................................................")

	length := 5
	breath := 9
	fmt.Printf("RECTANGLE ::::: Length: %v |Breath: %v | Area: %v | Perimeter: %v  \n", length, breath, 2*(length+breath), length*breath)
	fmt.Println(".......................................................................")

	hieght := 7
	fmt.Printf("ISO_TRIANGLE :::::: Heigth: %v |Breath: %v | Area: %v | Perimeter: %v  \n ", hieght, breath, half*hieght*breath, 2*hieght+breath)
	fmt.Println(".......................................................................")

	radius := 7
	fmt.Printf("CIRCLE :::: Radius: %v | Area: %v | Perimeter: %v  \n", radius, Pi*radius*radius, 2*Pi*radius)
	fmt.Println(".......................................................................")

	side1 := 7
	side2 := 8
	side3 := 9
	semi := (side1 + side2 + side3) / 2
	fmt.Printf("SCALENE_TRIANGLE :::: 1st Side: %v | 2nd Side: %v | 3rd side : %v | Semiperimeter: %v | Area: %v | Perimeter: %v  \n", side1, side2, side3, semi, (semi * (semi - side1) * (semi - side2) * (semi - side3)), side1+side2+side3)
	fmt.Println(".......................................................................")

	// 3D SHAPES
	fmt.Println(" \n ############################################## 3D SHAPE ######################################### \n ")
	fmt.Println("...............................................................................................................")

	fmt.Printf("CYLINDER :::: RADIUS: %v | HEIGHT :%v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", radius, hieght, 22/7*radius*radius*hieght, 2*(22/7)*radius*hieght+2*(22/7)*radius*radius)
	fmt.Println("...............................................................................................................")

	fmt.Printf("CONE :::: RADIUS: %v | HEIGHT :%v | LENGTH : %v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", radius, hieght, length, (1/3)*(22/7)*radius*radius*hieght, (22/7)*radius*(radius+length))
	fmt.Println("...............................................................................................................")

	fmt.Printf("CUBOID :::: LENGTH: %v | BREATH : %v | HEIGHT : %v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", length, breath, hieght, length*breath*hieght, 2*(length*breath+breath*hieght+hieght*breath))
	fmt.Println("...............................................................................................................")

	fmt.Printf("CUBOID :::: LENGTH: %v | BREATH : %v | HEIGHT : %v |  VOLUME : %v | TOTAL SURFACE AREA : %v \n ", length, breath, hieght, length*breath*hieght, 2*(length*breath+breath*hieght+hieght*breath))
	fmt.Println("...............................................................................................................")
}
