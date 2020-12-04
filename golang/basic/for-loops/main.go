package main

import "fmt"

func main() {

	// // basic
	// counter := 0
	// for counter < 5 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// // or without condition after for
	// counter := 0
	// for {
	// 	fmt.Println(counter)
	// 	counter++
	// 	if counter == 5 {
	// 		break
	// 	}
	// }

	// with statement
	for counter := 1; counter < 5; counter++ {
		fmt.Println(counter)
	}

}
