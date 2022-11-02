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

func main() {
	greeting("Rose")
	sayBye("Rose")
	loopNames([]string{"Rose", "Emma", "kate"}, greeting)
	loopNames([]string{"Rose", "Emma", "kate"}, sayBye)
	fmt.Print(math.Round(circle(12.12)))
}
