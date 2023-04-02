package app

import (
	"fmt"
	"go-learning/utils"
)

func Run() {
	arrayLearning()
	//sliceLearning()
	//mapLearning()
}

func arrayLearning() {
	//numbers := [5]int{5, 2, 4, 1, 3}
	//fmt.Println("array = ", numbers)
	//
	//fmt.Println("Printing evens only")
	//for _, num := range numbers {
	//	if num%2 == 0 {
	//		fmt.Println(num)
	//	}
	//}

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
	//numbers := []int{5, 2, 4, 1, 3}
	//fmt.Println("full slice = ", numbers)
	//first3 := numbers[1:4]
	//fmt.Println("slice partition = ", first3)
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
