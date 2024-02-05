package main

import "fmt"


type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	book.Pages = pages
}

func main() {

	Book1 := Book {
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages: 180,
	}
	Book2 := Book {
		Title: "To Kill a Mockingbird",
		Author: "Harper Lee",
		Pages: 281,
	}
	Book3 := Book {
		Title: "Pride and Prejudice",
		Author: "Jane Austen",
		Pages: 279,
	}

	updatePages(&Book1, 210)
	updatePages(&Book2, 250)
	updatePages(&Book3, 295)

        fmt.Println(Book1)
	fmt.Println(Book2)
	fmt.Println(Book3)

}

