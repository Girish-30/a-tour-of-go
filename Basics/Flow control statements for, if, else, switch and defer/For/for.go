package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for i = 0; i < 10; i++ {

		sum += i
		fmt.Println("Sonali", i, "Girish", sum)
	}
	scam := sum + i

	fmt.Println(sum)
	fmt.Println(i == sum)
	fmt.Println(scam)

}
