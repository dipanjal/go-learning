package main

import "fmt"

func main() {
	books := []string{"Java", "Python"}
	fmt.Println(books)
	fmt.Printf("Reference before append %p\n", &books)

	// append doesn't modify the slice,
	//instead it creates a new slice which we must to re-assign
	booksModified := append(books, "Go in Action")
	fmt.Println(booksModified)
	fmt.Printf("Reference after append %p\n", &booksModified)
}
