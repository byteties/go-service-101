package main


import (
	"fmt"
	"strings"
)
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func upperAllLetter(str *string) {
	*str = strings.ToUpper(*str)
}

func main() {
	i, j := 42, 2701

	fmt.Println("------------")
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	fmt.Println("------------")
	p = &j         // point to j
	*p = *p - 1   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println("------------")
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Println("------------")
	v.X = 4
	fmt.Println(v.X)

	fmt.Println("------------")
	k := &v
	k.X = 2000
	fmt.Println(v)

	fmt.Println("------------")
	fmt.Println(v1, p, v2, v3)

	fmt.Println("------------")
	name := "Weerasak"
	upperAllLetter(&name)
	fmt.Println(name)
}