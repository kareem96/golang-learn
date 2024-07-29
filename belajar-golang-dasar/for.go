package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("perulangna ke ", counter)
	// 	counter++
	// }

	// fmt.Println("selesai")

	//for with statement
	// init statement ->  counter := 1;
	// post statement ->  counter++
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	names := []string{"abdul", "kareem"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}

	//for range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}


	//break and continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Fizz")
			continue
		}
		fmt.Println("perulangan ke ", i)
	}

}
