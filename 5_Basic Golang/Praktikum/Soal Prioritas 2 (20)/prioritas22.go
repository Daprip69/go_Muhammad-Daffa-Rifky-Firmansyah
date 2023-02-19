//  Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka, seperti dibawah ini!

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Input Number: ")
	fmt.Scan(&number)

	for a := 1; a <= number; a++ {
		if number%a == 0 {
			fmt.Println("Faktorisasi dari  angka tersebut adalah:", a)
		}
	}

}

// -Output-
// PS C:\MSIB\4_Intro Algorithm & Basic Programming\Praktikum\Soal Prioritas 2 (20)> go run prioritas22.go
// Input Number: 69
// Faktorisasi dari  angka tersebut adalah: 1
// Faktorisasi dari  angka tersebut adalah: 3
// Faktorisasi dari  angka tersebut adalah: 23
// Faktorisasi dari  angka tersebut adalah: 69
