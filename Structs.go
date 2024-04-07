package main

import "fmt"

type book struct {
	name   string
	number int
}

func (books book) print_details() {
	fmt.Println(books.name)
	fmt.Println(books.number)
}

func main() {
	book1 := book{"A song of fire and ice", 1}
	book1.print_details()

}
