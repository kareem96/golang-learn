package main

import "fmt"

// variable argume [vararg]
func sumAll(numbers ...int) int{
	total := 0
	for _, numbers := range numbers{
		total += numbers
	}

	return total
}

// using slice
// func sumAll1(numbers []int) int{
// 	total := 0
// 	for _, numbers := range numbers{
// 		total += numbers
// 	}

// 	return total
// }

func main() {
	// fmt.Println(sumAll1([]int{10, 10, 10}))

	fmt.Println(sumAll(10, 10, 10))

	// else ex
	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

}

