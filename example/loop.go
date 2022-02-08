package main

import "fmt"


func main() {
	fmt.Println("----------------")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("----------------")
	sum2 := 1
	// for ; sum2 < 1000; {
	// 	sum2 += sum2
	// }
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)


	//last-in-first-out 
	fmt.Println("--------------------")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
