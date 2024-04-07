package main

import "fmt"

func main() {
	var map1 = make(map[int]int)
	fmt.Println(map1)

	map1[1] = 5
	map1[2] = 7

	fmt.Print(map1)
}
