package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println("------------")

	var sa []int = primes[1:4]
	fmt.Println(sa)

	fmt.Println("------------")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	fmt.Println("------------")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	fmt.Println("------------")
	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:2]
	fmt.Println(s1)

	s1 = s1[1:]
	fmt.Println(s1)

	s1 = s1[1:4]
	fmt.Println(s1)

	fmt.Println("------------")
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)
	// Slice the slice to give it zero length.
	s2 = s2[:0]
	printSlice(s2)

	s2 = s2[:6]
	printSlice(s2)
	
	// Extend its length.
	s2 = s2[:4]
	printSlice(s2)

	// Drop its first two values.
	s2 = s2[2:]
	printSlice(s2)

	s2 = s2[:4]
	printSlice(s2)

	s2 = s2[1:]
	printSlice(s2)
	
	s2 = s2[1:]
	printSlice(s2)
}
