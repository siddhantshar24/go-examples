package main

import "fmt"

type sport interface {
	sportName() string
}
type human struct {
	name  string
	sport string
}

func (h human) sportName() string {
	return h.name + " " + h.sport
}

func main() {
	human1 := human{"Siddhant", "basketball"}
	fmt.Println(human1.sportName())

}
