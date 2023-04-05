package main

import "fmt"

func main() {
	//example1()
	bookSliceTypeExample()
}

func example1() {
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

/**
in this method we are going to create a custom type of string slice, Books
which is a slice of string, and we can use it same as a slice of string
*/
func bookSliceTypeExample() {
	type Books []string
	books := Books{"Java", "Python", "Go", "Rust"}
	fmt.Println(books)

	for index, book := range books {
		fmt.Println(index, "->", book)
	}

	// calling the generic function which can accept any type which is extended from as []string
	printSlice(books, false)
}

// This function accepts any type which is extended from []string
func printSlice(items []string, showIndex bool) {
	for index, item := range items {
		if showIndex {
			fmt.Println(index, "->", item)
		} else {
			fmt.Println(item)
		}
	}
}
