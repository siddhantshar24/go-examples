package main

import "fmt"

func main() {
	map1 := map[int]int{
		1: 5,
		2: 7,
	}
	fmt.Println(map1)

	for id, value := range map1 {
		fmt.Println(id, value)
	}
}
