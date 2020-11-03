package main

import "fmt"

func main() {
	// Hello World
	// fmt.Println("Hello World")

	// Looping for - range
	// var os = [3]string{"linux", "windows", "macos"}

	// for i, e := range os {
	// 	fmt.Println(i, e)
	// }

	// Multidimentional array
	// var numbers1 = [2][3]int{
	// 	[3]int{1, 2, 3},
	// 	[3]int{5, 6, 7},
	// }

	// var numbers2 = [2][3]int{
	// 	{3, 2, 3},
	// 	{3, 4, 5},
	// }

	// fmt.Println(numbers1)
	// fmt.Println(numbers2)

	// defer

	// defer fmt.Println("halo")
	// fmt.Println("selamat datang")

	// slicing teknik 3 index
	// var fruits = []string{"apple", "grape", "banana"}
	// var bFruits = fruits[0:2:2]
	// // var cFruits = fruits[0:2:1]

	// fmt.Println(bFruits)
	// fmt.Println(cap(bFruits))

	// invoking func printMessage()
	// var names = []string{"hpazk", "zkh"}
	// printMessage("hai", names)

	// invoking func randomValue()
	// rand.Seed(time.Now().Unix())
	// var val1 = randomValue(1, 3)
	// var val2 = randomValue(1, 3)
	// fmt.Println(val1)
	// fmt.Println(val2)

	// invoking function return keyword
	// divideNumber(10, 2)
	// divideNumber(4, 0)
	// divideNumber(8, -4)

	// invoking function hitung()
	// diameterLingkaran := 2.5
	// var luas, keliling = hitung(diameterLingkaran)

	// fmt.Printf("luas= %.3f dan keliling= %.2f\n", luas, keliling)

	// invoking function hitung() - predefined return value
	// diameter := 4.5
	// var luas, keliling = hitung(diameter)

	// fmt.Printf("luas= %.3f dan keliling= %.2f\n", luas, keliling)

	// slice as argument in variadic function
	// var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	// var avg = calculate(numbers...)
	// var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	// fmt.Println(msg)

	// invoking varidic function
	// var yourName, yourHobbies = profile("hpazk", "coding", "table tennis")

	// fmt.Println(yourName, yourHobbies)

	// var hobbies = []string{"coding", "table tennis"}

	// var yourName, yourHobbies = profile("hpazk", hobbies...)
	// fmt.Println(yourName, yourHobbies)

	// var getMinMax = func(nums []int) (int, int) {
	// 	var min, max int

	// 	for i, e := range nums {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}
	// 	return min, max
	// }

	// var numbers = []int{0, 1, 2, 3, 4, 5, 6}
	// var min, max = getMinMax(numbers)

	// fmt.Println(min, max)

	// IIFE closure
	// var numbers = []int{1, 2, 3, 2, 3, 4, 5, 6, 5, 7}

	// var newNumber = func(min int) []int {
	// 	var r []int
	// 	for _, e := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 	}
	// 	return r
	// }(4)

	// fmt.Println("numbers: ", numbers)
	// fmt.Println("new numbers: ", newNumber)

	// closure with manifest typing style
	// var list = []string{"satu", "dua", "tiga"}

	// var closure (func(string, int, []string) string) = func(a string, b int, c []string) string {
	// 	text := fmt.Sprintf("val a = %s, val b = %d, val c = %v", a, b, strings.Join(c, ","))
	// 	return text
	// }

	// // Or use variable assignment
	// // closure = func(a string, b int, c []string) string {
	// // 	text := fmt.Sprintf("val a = %s, val b = %d, val c = %v", a, b, strings.Join(c, ","))
	// // 	return text
	// // }

	// fmt.Println(closure("tes", 123, list))

	// pointer
	// var nilaiA = 4
	// var nilaiB = &nilaiA

	// fmt.Println("nilaiA", nilaiA)
	// fmt.Println("alamat memory nilaiA", &nilaiA)

	// fmt.Println("nilaiB", *nilaiB)
	// fmt.Println("alamat memory nilaiB", nilaiB)

	// // instance the student struct
	// var student1 Student
	// student1.name = "hpazk"
	// student1.grade = "A"

	// // shorthand instace struct object
	// var student2 = Student{"zkh", "B"}
	// // with key
	// var student3 = Student{name: "ltg"}
	// student3.grade = "A"

	// // fmt.Println(student1)
	// // fmt.Println(student2)
	// // fmt.Println(student3)

	// // variable object pointer
	// var std1 *Student = &student2
	// fmt.Println(std1)
	// fmt.Println("name", std1.name)
	// fmt.Println("grade", std1.grade)

	// std1.name = "hpazk"
	// std1.grade = "A"

	// fmt.Println("name after re-assign:", std1.name)
	// fmt.Println("grade after re-assign:", std1.grade)

	// Embedded Struct Implementation
	// var std = Student{}
	// std.name = "hpazk"
	// std.age = 26
	// std.address = "sumedang"
	// std.classroom = "A"
	// std.grade = "A"
	// std.mark = 90

	// fmt.Println("Student:", std)
	// fmt.Println("Student Name:", std.name)
	// fmt.Println("Student Class Room:", std.classroom)
	// fmt.Println("Student Mark:", std.mark)

	var std1 = Student{}
	std1.Person.name = "ltg"
	std1.classroom = "B"
	std1.mark = 85
	std1.Person.age = 26
	std1.age = 25

	fmt.Println("Student:", std1)
	fmt.Println("Student Name:", std1.name)
	fmt.Println("Student Class Room:", std1.classroom)
	fmt.Println("Student Mark:", std1.mark)
	fmt.Println("Student Age in Person struct:", std1.age)
	fmt.Println("Student Age in Student struct:", std1.Person.age)
}

// function without return type
// func printMessage(message string, persons []string) {
// 	var name = strings.Join(persons, ", ")
// 	fmt.Println(message, name)
// }

// function with return type
// func randomValue(min, max int) int {
// 	var value = rand.Int()%(max-min+1) + min
// 	return value
// }

// breakup function with return keyword
// func divideNumber(m, n int) {
// 	if n == 0 {
// 		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// 		return
// 	}
// 	var res = m / n
// 	fmt.Printf("%d / %d = %d\n", m, n, res)
// }

// function multiple return value
// func hitung(diameter float64) (float64, float64) {
// 	var luas = math.Pi * math.Pow(diameter/2, 2)

// 	var keliling = math.Pi * diameter

// 	return luas, keliling
// }

// function predefined return value
// func hitung(diameter float64) (luas float64, keliling float64) {
// 	luas = math.Pi * math.Pow(diameter/2, 2)

// 	keliling = math.Pi * diameter

// 	return
// }

// variadic function with slice as parameter
// func calculate(numbers ...int) float64 {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	var avg = float64(total) / float64(len(numbers))
// 	return avg
// }

// func profile(name string, hobbies ...string) (string, string) {
// 	var yourName = name
// 	var yourHobbies = strings.Join(hobbies, ", ")

// 	return yourName, yourHobbies
// }

/*
Struct must add comment like below:

This linter error is caused by your Agent type being exported, even if its attributes are not. To not export your type, define it in lowercase like such:

type student struct {
	id    int
	name  string
	grade string
}

The reason why your linter complains about this is that godoc uses those comments to automatically generate documentation for your projects. You can find many examples of such documented Go projects at pkg.go.dev.

If you upload one of your Go projects to GitHub for example, pkg.go.dev will automatically generate a documentation for you using those comments. You can even add runnable code examples and many other things, as shown on go-doc tricks.
*/

// Basic Struct

// Student is...
// type Student struct {
// 	name  string
// 	grade string
// }

// Embedded Struct

// Person is..
type Person struct {
	name    string
	age     int
	address string
}

// Student is..
type Student struct {
	Person
	age       int
	classroom string
	grade     string
	mark      int
}
