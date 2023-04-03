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

	// print a slice using for loop,
	// here we are going to get a tuple like
	// the first one is the index of the item and
	// the second one is the item inside the slice
	for index, item := range booksModified {
		fmt.Println(index, "->", item)
	}

	// but what if we don't want to use the index
	// we can just ignore it using low dash like this _
	for _, item := range booksModified {
		fmt.Println(item)
	}
}
