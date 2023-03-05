// Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut receiver kepada struct Car yang memiliki property type dan fuelIn 

package main

import "fmt"

type Car struct {
	Type string
	FuelIn float64
}
func (e Car) EstimatedRange() float64 {
	return e.FuelIn / 1.5
}
func main() {
	var Car0 Car
	Car0.Type = "Sedan"
	Car0.FuelIn = 69
	
	fmt.Println("Dengan bahan bakar sebanyak:", Car0.FuelIn, ";Maka mobil dengan tipe: ", Car0.Type, ";Dapat menempuh jarak sejauh", Car0.EstimatedRange(), "mill")
}