package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println("---------------------")

	// var k = map[string]Vertex{
	// 	"test":Vertex{
	// 		40,50,
	// 	},
	// 	"data":Vertex{
	// 		10,20,
	// 	},
	// }

	// fmt.Println(k)

	// type Vertex struct {
	// 	Lat, Long float64
	// }
	
	var j = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(j)

	fmt.Println("---------------------")
	k := make(map[string]int)

	k["Answer"] = 42
	fmt.Println("The value:", k["Answer"])

	k["Answer"] = 48
	fmt.Println("The value:", k["Answer"])

	delete(k, "Answer")
	fmt.Println("The value:", k["Answer"])

	v, ok := k["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
