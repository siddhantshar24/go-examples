package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	var fact int
	fact = 1

	for i := 1; i <= a; i++ {
		fact = fact * i
	}

	fmt.Println(fact)

}
