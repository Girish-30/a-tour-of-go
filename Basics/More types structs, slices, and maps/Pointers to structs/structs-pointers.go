package main

import "fmt"

type Vertex struct {
	x, y, z, a, b, c int
}

func main() {
	v := Vertex{1, 2, 5, 6, 7, 8}
	p := &v
	p.x = 19
	p.z = 1e6
	p.b = 4e6

	fmt.Println(v)
}
