package main

import "fmt"

func main() {
	map1 := map[int]int{
		1: 9,
	}
	fmt.Println(map1)
	delete(map1, 1)
	fmt.Println(map1)

}
