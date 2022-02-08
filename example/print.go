package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)
func add(x,y int) int {
	return x + y
}

func swap(x,y string) (string,string) {
	return y , x
}

func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return  
}
var c, python, java bool
var i, j int = 1, 2

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("-------------------------")
	fmt.Println("The time is", time.Now())

	//สภาพแวดล้อมที่โปรแกรมเหล่านี้ทำงานนั้น ถูกกำหนดขึ้นอย่างตั้งใจ นั่นทำให้แต่ละครั้งที่โปรแกรมตัวอย่างนี้ทำงาน rand.Intn จะคืนเลขเดิมมาเสมอ
	//(ถ้าอยากเห็นค่าอื่น ต้องมีตัวสร้าง seed; ดูได้ที่ rand.Seed เวลาใน playground ก็เป็นค่าคงที่ ดังนั้นคุณต้องหาอะไรอย่างอื่นมาทำ seed แทนด้วย)
	fmt.Println("-------------------------")
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("-------------------------")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	//การตั้งชื่อมีความสำคัญ ถ้าชื่อขึ้นต้นด้วยตัวพิมพ์ใหญ่ ตัวอย่างเช่น Pizza จะถือว่าเป็นชื่อสามารถถูกใช้จากภายนอกได้ เหมือนกับ Pi ที่ถูกส่งออกมาจากแพ็คเกจ math
	fmt.Println("-------------------------")
	fmt.Println(math.Pi)

	fmt.Println("-------------------------")
	fmt.Println(add(42, 13))

	fmt.Println("-------------------------")
	a, b := swap("hello", "world")
	fmt.Println(swap(b,a))

	fmt.Println("-------------------------")
	fmt.Println(split(17))

	fmt.Println("-------------------------")
	var i int
	fmt.Println(i, c, python, java)

	fmt.Println("-------------------------")
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Println("-------------------------")
	k := 3
	fmt.Println(i, j, k, c, python, java)
}
