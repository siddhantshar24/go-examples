package main

import "fmt"

func main() {
	map1 := map[int]int{
		1: 9,
	}
	value1, ok := map1[2]
	fmt.Println(value1, ok)
	value2, ok := map1[1]
	fmt.Println(value2, ok)
}
