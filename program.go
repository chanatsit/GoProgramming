package main

import "fmt"

func main() {
	// เปลี่ยนแปลง
	age := 80
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
	fmt.Println("ผลบวก  = ", number1+number2)
	fmt.Println("ผลหาร = ", number1/number2)
	fmt.Println("เศษผลหาร = ", number1%number2)
	fmt.Printf("ค่าเฉลี่ยที่ได้คือ %v\n", (number1+number2)/2)
	// ตัวดำเนินการเปรียบเทียบ ==, !=, >, <, >=, <=
	fmt.Println("จำนวนสองจำนวนนี้เท่ากันหรือไม่ ", number1 == number2)
	fmt.Println("จำนวนที่หนึ่งมากกว่าจำนวนที่สองหรือไม่ ", number1 > number2)

}
