package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("---------------------")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("---------------------")
	var x, y int = 3, 4
	// var g float64 = math.Sqrt(float64(x*x + y*y))
	g := math.Sqrt(float64(x*x + y*y))
	// var k uint = uint(g)
	k := uint(g)
	fmt.Println(x, y, k)

	fmt.Println("---------------------")
	v := 3.142    // 42 , 0.867 + 0.5i 
	fmt.Printf("v is of type %T\n", v)

	fmt.Println("---------------------")
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	fmt.Println("---------------------")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big)) error
}


//bool

//string

//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr

//byte // alias for uint8

//rune // alias for int32
     // represents a Unicode code point

//float32 float64

//complex64 complex128