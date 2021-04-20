package main

import "fmt"

func main() {

	var triangleBase float64
	var triangleHeight float64
	var areaOfTriangle float64
	fmt.Println("Enter the base of triangle:")
	fmt.Scanf("%f", &triangleBase)
	fmt.Println("Enter the height of triangle:")
	fmt.Scanf("%f", &triangleHeight)
	areaOfTriangle = (triangleBase * triangleHeight) / 2
	fmt.Println("Area of triangle is: ", areaOfTriangle)
}
