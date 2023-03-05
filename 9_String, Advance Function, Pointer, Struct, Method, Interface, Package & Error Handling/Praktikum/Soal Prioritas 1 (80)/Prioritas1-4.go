// Buatlah program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!
package main

import "fmt"

func getMinMax(numbers ...*int) (minim int, maxim int) {
	minim = *numbers[0]
	maxim = *numbers[0]

	for _, num := range numbers {
		if *num < minim {
			minim= *num
		}
		if *num > maxim {
			maxim = *num
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, minim, maxim int
	fmt.Println("Please enter numbers 6 times")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	minim, maxim = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("%d adalah nilai minim", minim)
	fmt.Printf("%d adalah nilai maxim", maxim)
}
