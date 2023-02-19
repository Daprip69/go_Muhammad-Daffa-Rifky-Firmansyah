// buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilang ganjil atau genap

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Input Number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("Bilangan %d adalah genap", num)
	} else {
		fmt.Printf("Bilangan %d adalah ganjil", num)
	}
}
