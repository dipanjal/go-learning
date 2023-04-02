package app

import (
	"fmt"
	"go-learning/utils"
)

func Run() {
	arrayLearning()
	fmt.Println()
	sliceLearning()
	fmt.Println()
	mapLearning()
}

func arrayLearning() {
	fmt.Println("Array Learning....")
	bookShelf := [5]string{
		"Pother Pachalli",
		"Harry Potter",
		"Java Learning",
		"Python Learning",
		"Go In Action",
	}
	fmt.Println("Looping through range keyword....")
	for _, item := range bookShelf {
		fmt.Println(item)
	}
	fmt.Println("Manual For Loop....")
	for i := 0; i < len(bookShelf); i++ {
		fmt.Println(bookShelf[i])
	}

}

func sliceLearning() {
	fmt.Println("Learning Slice.....")
	books := []string{
		"Pother Pachalli",
		"Harry Potter",
		"Java Learning",
		"Python Learning",
		"Go In Action",
	}
	fmt.Println("Looping through range keyword....")
	for _, book := range books {
		fmt.Println(book)
	}
	fmt.Println("Looping manually....")
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
}

func mapLearning() {
	myMap := map[string]float32{
		"dip":    100.00,
		"shoaib": 200.00,
	}
	key := "fahim"
	salary, err := utils.GetValueFromMap(myMap, key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(key, "has salary", salary)
	}
}
