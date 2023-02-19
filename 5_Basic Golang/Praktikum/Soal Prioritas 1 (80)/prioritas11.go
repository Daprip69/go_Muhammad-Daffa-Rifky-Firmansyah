// buatlah sebuah program untuk menghitung luas trapesium

package main

import "fmt"

func main() {
	var a, t, s float64

	fmt.Print("Input Sisi Atas Trapesium: ")
	fmt.Scan(&a)

	fmt.Print("Input Sisi Bawah Trapesium: ")
	fmt.Scan(&t)

	fmt.Print("Input Sisi Miring Trapesium: ")
	fmt.Scan(&s)

	luas := (a + s) * t / 2

	fmt.Printf("Hasil dari Luas Trapesium: %2f", luas)
}
