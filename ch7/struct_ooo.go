package main

import "fmt"

type Book struct {
	name   string
	author string
	pages  int
}

func (book Book) print_details() {
	fmt.Printf("Book %s was written by %s", book.name, book.author)
	fmt.Printf("\nIt contains %d pages.\n", book.pages)
}

func main() {
	// declaring a struct instance
    book1 := Book{"Monster Blood", "R.L.Stine", 131}
	book1.print_details()

	book1.name = "Vampire Breath"
    book1.pages = 162
	book1.print_details()
}