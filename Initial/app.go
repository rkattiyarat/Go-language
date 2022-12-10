package main

import (
	"fmt"
	"math"
)

func greeting(name string) {
	fmt.Printf("Hello %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)

}

func loopNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func circle(radius float64) float64 {
	return math.Pi * (radius * radius)
}

//this function use var from test.go variable which is outside main func but in the same package and call sayAge in main func
func sayAge(age int) {
	fmt.Printf("I am %v years old \n", age)
}
