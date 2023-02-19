package main

import "fmt"

func main() {
	triangle()
	fmt.Print("Segitiga Triangle")
}

func triangle() {
	for a := 1; a <= 5; a++ {
		for b := 5; b >= a; b-- {
			fmt.Print(" ")
		}
		for c := 1; c <= a; c++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

