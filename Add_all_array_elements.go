package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var z int = 0
	for i := 0; i <= 4; i++ {
		z += arr[i]
	}
	fmt.Println(z)
}
