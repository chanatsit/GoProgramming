package main

import "fmt"

func main() {
	// เปลี่ยนแปลง
	age := 80
	//var name1 string
	//var name string = "Chanat"
	name := "chanat"
	fmt.Println("Hello Programming")
	fmt.Println("Chanat  Sittipolpat")
	fmt.Println("The age is ", age)
	fmt.Println("name is ", name)
	fmt.Printf("Type of age is %T", age)
	// ตัวดำเนินการทางคณิตศาสตร์
	// var number1 int = 10
	// var number1, number2 = 10, 3
	// number1 , number2 := 10, 3
	number1 := 10
	number2 := 3
	var score int
	fmt.Println("ผลบวก  = ", number1+number2)
	fmt.Println("ผลหาร = ", number1/number2)
	fmt.Println("เศษผลหาร = ", number1%number2)
	fmt.Printf("ค่าเฉลี่ยที่ได้คือ %v\n", (number1+number2)/2)
	// ตัวดำเนินการเปรียบเทียบ ==, !=, >, <, >=, <=
	fmt.Println("จำนวนสองจำนวนนี้เท่ากันหรือไม่ ", number1 == number2)
	fmt.Println("จำนวนที่หนึ่งมากกว่าจำนวนที่สองหรือไม่ ", number1 > number2)

	//fmt.Print("ชื่อนักเรียน : ")
	//fmt.Scanf("%s", &name1)
	fmt.Print("คะแนนที่ได้รับ : ")
	fmt.Scanf("%d", &score)
	//fmt.Println("สวัสดี ", name1)
	fmt.Println("คะแนนที่ได้ ", score)
	if score >= 60 {
		fmt.Println("คุณผ่านการทดสอบ")
	} else {
		fmt.Println("คุณไม่ผ่านการทดสอบ")
	}
	number3 := 0
	switch number3 {
	case 1:
		fmt.Println("hello World")
	case 2:
		fmt.Println("Hello, Thailand")
	default:
		fmt.Println("ไม่ถูกต้อง")
	}
	// array
	//var number5 [3]int = [3]int{100, 200, 300}
	number4 := [3]int{100, 200, 300}
	fmt.Println(number4)
	fmt.Println(len(number4))
	number5 := [...]int{100, 200, 300, 400} // เพิ่มข้อมูลได้เรื่อยๆ
	fmt.Println(len(number5))
	fmt.Println(number5[0])
	number6 := [3]int{10, 20} // default => [10,20,0]
	fmt.Println(number6)
}
