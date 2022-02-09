package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 2)
}

func intSeq() func() int {
	i := 0
	return func() int {
			i++
			return i
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	n := 0
	return func() int {
		if n++; n == 1 {
			return 0
		}
		a, b = b, a + b
		return a
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return x*x + y*y //math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("--------------------")
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println("--------------------")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
