package main

import "fmt"

// type myInfo interface {
// 	details() string
// }

// type Name struct {
// 	fName, lName string
// }

// type Hobby struct {
// 	title, kind string
// }

// func (n Name) details() string {
// 	return n.fName + " " + n.lName
// }

// func (h Hobby) details() string {
// 	return h.title + ": " + h.kind
// }

// func getMyInfo(info myInfo) string {
// 	return info.details()
// }

// func multiVals() (string, int, bool) {
// 	return "Rose", 1995, true
// }

// func sum(numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n

// 	}
// 	return total
// }

func greeting() {
	fmt.Println("Hello from a regular function")
}

func anotherFunc(f func()) {
	fmt.Println("Say Hello from function that passed as an argument")
}
func main() {
	greeting()
	anotherFunc(greeting)

	// 	myAge := 25

	// 	addOneYearToAge := func() int {
	// 		i := 1
	// 		return myAge + i

	// 	}
	// 	fmt.Println(addOneYearToAge())

	// 	sumNumbers := sum(2, 4, 6, 8, 10)
	// 	fmt.Println(sumNumbers)
	// 	a, b, c := multiVals()
	// 	fmt.Printf("%v was born in %v. Is this %v?\n", a, b, c)

	// 	names := Name{fName: "Rose", lName: "Kattiyarat"}
	// 	hobbies := Hobby{title: "Swimming", kind: "Sport"}

	// 	fmt.Println(names)
	// 	fmt.Println(hobbies)
	// 	fmt.Println(getMyInfo(names), (hobbies))

	// 	number := 10
	// 	var n = &number
	// 	fmt.Println("Memory address is ", n)
	// 	fmt.Println("Value is ", *n)

	// 	var x, y, z int = 1, 2, 3
	// 	fmt.Println(x, y, z)

	// var input string
	// fmt.Println(">>")
	// fmt.Scan(&input)
	// switch s := strings.ToLower(input); {
	// case s == "Morning":
	// 	fmt.Println("Breakfast?")
	// case s == "Afternoon":
	// 	fmt.Println("Lunch Time!")
	// case s == "Evening":
	// 	fmt.Println("Dinner is ready :)")
	// default:
	// 	fmt.Println("Good night Zzzz")
	// }

	// var input string
	// fmt.Println(">>")
	// fmt.Scan(&input)
	// switch input {
	// case "Morning":
	// 	fmt.Println("Breakfast?")
	// case "Afternoon":
	// 	fmt.Println("Lunch Time!")
	// case "Evening":
	// 	fmt.Println("Dinner is ready :)")
	// default:
	// 	fmt.Println("Good night Zzzz")
	// }

	// testNum := 20
	// fmt.Println(testNum)

	// newNum := &testNum
	// *newNum = 50
	// fmt.Println(&newNum)

	// myName := "Rose"
	// fmt.Println(myName)

	// changeName := &myName
	// *changeName = "Rungnapa"
	// fmt.Println(myName)

}

// func addNum()

// func a(f func(int) int) {
// 	fmt.Println(f(2))
// }

// func main() {
// 	name := "Rose"
// 	var name string = "Rungnapa"
// 	var y = 2
// 	var x = 5
// 	a(func(x) { return x + y })
// 	fmt.Println(name)
// }
