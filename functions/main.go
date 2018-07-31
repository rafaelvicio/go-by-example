package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func main() {

	rest := plus(1, 2)
	fmt.Println("Rest: ", rest)

	restPlus := plusPlus(1, 2, 3)
	fmt.Println("restPlus: ", restPlus)

	a, b := vals()
	fmt.Println("Vals: ", a, b)

}
